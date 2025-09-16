package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	embedServer "github.com/Kovarniykrab/serverTesting"
	"github.com/Kovarniykrab/serverTesting/api/handlers"
	"github.com/Kovarniykrab/serverTesting/api/routers"
	"github.com/Kovarniykrab/serverTesting/configs"
	"github.com/Kovarniykrab/serverTesting/database"
	_ "github.com/Kovarniykrab/serverTesting/docs"
	"github.com/pressly/goose/v3"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/valyala/fasthttp"
)

// @title          TestUser API
// @version        0.5
// @description    API для управления пользователями
// @host
// @BasePath       /
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

func main() {
	var _ = handlers.RegisterUserHandler
	fmt.Println("API server started on :8080")
	r := routers.GetRouter()

	db, err := database.DBInit()
	if err != nil {
		slog.Error("Database connection failed: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	certDirectory := "/etc/letsencrypt/live/wednode.ru"
	certFile := filepath.Join(certDirectory, "fullchain.pem")
	keyFile := filepath.Join(certDirectory, "privkey.pem")

	_, certErr := os.Stat(certFile)
	_, keyErr := os.Stat(keyFile)

	if certErr == nil && keyErr == nil {
		fmt.Printf("SSL found. Starting HTTPS server on :8080")
		err := fasthttp.ListenAndServeTLS(":8080", certFile, keyFile, r.Handler)
		if err != nil {
			slog.Error("HTTPS server failed", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("SSL certificates NOT FOUND. Starting HTTP server on :8080\n")
		err := fasthttp.ListenAndServe(":8080", r.Handler)
		if err != nil {
			slog.Error("HTTP server failed: %v", err)
			os.Exit(1)
		}
	}

}

func initLogger(level *slog.Level) *slog.Logger {
	var logLevel slog.Level

	if level == nil {
		logLevel = slog.LevelInfo
	} else {
		logLevel = *level
	}
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))

	return log
}

var embedMigrations embed.FS

func migrate(cfg configs.Config) {
	goose.SetBaseFS(embedServer.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.PSQL.DSN)))

	if err := goose.Up(db, "resources/store/psql/migrations"); err != nil {
		panic(err)
	}
}
