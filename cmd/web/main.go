package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/Broderick-Westrope/eventurely/internal/models"
	_ "github.com/mattn/go-sqlite3"
	flag "github.com/spf13/pflag"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type application struct {
	logger *slog.Logger
	events models.EventRepository
}

func main() {
	addr := flag.String("addr", ":2000", "HTTP network address")
	dsn := flag.String("dsn", "file:test.db", "SQLite data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger: logger,
		events: models.NewEventRepository(db),
	}

	app.logger.Info("starting server", slog.String("addr", *addr))

	err = http.ListenAndServe(
		*addr,
		h2c.NewHandler(app.routes(), &http2.Server{}),
	)
	app.logger.Error(err.Error())
	os.Exit(1)
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
