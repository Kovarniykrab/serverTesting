package database

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/Kovarniykrab/serverTesting/configs"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var DB *bun.DB

func New(cfg configs.PSQL) (*bun.DB, error) {
	dsn := cfg.DSN
	if dsn == "" {
		slog.Error("SERVER_PSQL_DSN environment variable is required")
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	slog.Info("Successfully connected to database")

	DB = db

	return db, nil
}
