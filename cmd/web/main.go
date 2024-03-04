package main

import (
	"context"
	"crypto/tls"
	"log/slog"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"github.com/Broderick-Westrope/eventurely/internal/models"
	"github.com/jackc/pgx/v5"
	_ "github.com/mattn/go-sqlite3"
	flag "github.com/spf13/pflag"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type application struct {
	logger      *slog.Logger
	events      models.EventRepository
	invitations models.InvitationRepository
}

func main() {
	flag.Usage = func() {
		_, _ = os.Stderr.WriteString("Usage: eventurely [flags]\n")
		flag.PrintDefaults()
	}
	addr := flag.String("addr", ":2000", "HTTP network address")
	dsn := flag.String("dsn", "postgresql://username:password@host:5432/database?sslmode=require", "PostgreSQL data source name")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, *dsn)
	if err != nil {
		logger.Error("unable to connect to database", slog.AnyValue(err))
		os.Exit(1)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err = conn.Close(ctx)
		if err != nil {
			logger.Error("unable to close database connection", slog.AnyValue(err))
			os.Exit(1)
		}
	}(conn, ctx)

	app := &application{
		logger:      logger,
		events:      models.NewEventRepository(conn),
		invitations: models.NewInvitationRepository(conn),
	}

	app.logger.Info("starting server", slog.String("addr", *addr))

	opts := connect.WithInterceptors(
		app.loggingInterceptor(),
	)

	var server *http.Server
	if *enableTLS {
		var tlsConfig *tls.Config
		tlsConfig, err = loadTLSConfig()
		if err != nil {
			logger.Error("failed to load TLS credentials", slog.AnyValue(err))
			os.Exit(1)
		}

		server = &http.Server{
			Addr:      *addr,
			Handler:   h2c.NewHandler(app.routes(opts), &http2.Server{}),
			TLSConfig: tlsConfig,
		}

		app.logger.Info("TLS is enabled")
		err = server.ListenAndServeTLS("", "") // The paths are already provided in TLSConfig
	} else {
		server = &http.Server{
			Addr:    *addr,
			Handler: h2c.NewHandler(app.routes(opts), &http2.Server{}),
		}

		err = server.ListenAndServe()
	}

	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}

func loadTLSConfig() (*tls.Config, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return config, nil
}
