package db

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

var (
	DB_TYPE     string
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_SSLMODE  string
)

func init() {
	DB_TYPE = "postgres"
	DB_NAME = "postgres"
	DB_USERNAME = "postgres"
	DB_PASSWORD = "postgres"
	DB_HOST = "db"
	DB_SSLMODE = "disable"
}

func Connect() *sql.DB {
	dsn := DB_TYPE + "://" + DB_USERNAME + ":" + DB_PASSWORD + "@" + DB_HOST + "/" + DB_NAME + "?sslmode=" + DB_SSLMODE

	slog.Info("Connecting to database")
	slog.Info("DSN: " + dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		slog.Error("failed to connect database", err)
	}

	if err = db.Ping(); err != nil {
		slog.Error("failed to ping database", err)
	}

	return db
}
