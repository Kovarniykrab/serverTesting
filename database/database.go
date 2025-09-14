package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var DB *bun.DB

func DBInit() (*bun.DB, error) {
	dsn := os.Getenv("SERVER_PSQL_DSN")
	if dsn == "" {
		log.Fatal("SERVER_PSQL_DSN environment variable is required")
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to database")

	return DB, nil
}
