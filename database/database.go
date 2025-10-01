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

type Repository struct {
	log    *slog.Logger
	db     *bun.DB
	ctx    context.Context
	cancel context.CancelFunc
}

func New(ctx context.Context, cfg configs.PSQL, log *slog.Logger) (*Repository, error) {
	dsn := cfg.DSN
	if dsn == "" {
		log.Error("DSN environment variable is required")
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	ctx, cancel := context.WithCancel(context.Background())
	if err := db.PingContext(ctx); err != nil {
		cancel()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Info("Successfully connected to database")

	service := &Repository{
		log:    log,
		db:     db,
		ctx:    ctx,
		cancel: cancel,
	}

	return service, nil
}

func (s *Repository) Close() error {
	if s.cancel != nil {
		s.cancel()
	}

	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
