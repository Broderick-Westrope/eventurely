package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	flag "github.com/spf13/pflag"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":2000", "HTTP network address")
	dsn := flag.String("dsn", "file:test.db?cache=shared&mode=memory", "SQLite data source name")
	flag.Parse()

	app := &application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	db, err := openDB(*dsn)
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app.logger.Info("starting server", slog.String("addr", *addr))

	err = http.ListenAndServe(*addr, app.routes())
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
