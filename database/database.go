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

type Repository struct {
	log *slog.Logger
	db  *bun.DB
}

func New(cfg configs.PSQL, log *slog.Logger) (*Repository, error) {
	dsn := cfg.DSN
	if dsn == "" {
		log.Error("DSN environment variable is required")
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Info("Successfully connected to database")

	service := &Repository{
		log: log,
		db:  db,
	}

	return service, nil
}

func (s *Repository) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
