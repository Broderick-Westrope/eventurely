package main

import (
	"crypto/tls"
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"github.com/Broderick-Westrope/eventurely/internal/models"
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
	addr := flag.String("addr", ":2000", "HTTP network address")
	dsn := flag.String("dsn", "file:test.db", "SQLite data source name")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger:      logger,
		events:      models.NewEventRepository(db),
		invitations: models.NewInvitationRepository(db),
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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
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
