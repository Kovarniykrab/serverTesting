package main

import (
	"database/sql"
	"embed"
	"log/slog"
	"os"
	"path/filepath"

	embedServer "github.com/Kovarniykrab/serverTesting"
	"github.com/Kovarniykrab/serverTesting/api/handlers"
	"github.com/Kovarniykrab/serverTesting/api/routers"
	"github.com/Kovarniykrab/serverTesting/configs"
	"github.com/Kovarniykrab/serverTesting/database"
	_ "github.com/Kovarniykrab/serverTesting/docs"
	"github.com/jessevdk/go-flags"
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

	conf := configs.Config{}
	parser := flags.NewParser(&conf, flags.Default)
	if _, err := parser.Parse(); err != nil {
		panic(err)
	}

	var _ = handlers.RegisterUserHandler
	slog.Info("API server started on :8080")
	r := routers.GetRouter()

	db, err := database.New(conf.PSQL)
	if err != nil {
		slog.Error("Database connection failed: %v", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	migrate(conf.PSQL)

	certDirectory := "/etc/letsencrypt/live/wednode.ru"
	certFile := filepath.Join(certDirectory, "fullchain.pem")
	keyFile := filepath.Join(certDirectory, "privkey.pem")

	_, certErr := os.Stat(certFile)
	_, keyErr := os.Stat(keyFile)

	if certErr == nil && keyErr == nil {
		slog.Info("SSL found. Starting HTTPS server on :8080")
		err := fasthttp.ListenAndServeTLS(":8080", certFile, keyFile, r.Handler)
		if err != nil {
			slog.Error("HTTPS server failed", "error", err)
			os.Exit(1)
		}
	} else {
		slog.Info("SSL certificates NOT FOUND. Starting HTTP server on :8080\n")
		err := fasthttp.ListenAndServe(":8080", r.Handler)
		if err != nil {
			slog.Error("HTTP server failed: %v", "error", err)
			os.Exit(1)
		}
	}

}

var embedMigrations embed.FS

func migrate(cfg configs.PSQL) {
	goose.SetBaseFS(embedServer.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.DSN)))

	if err := goose.Up(db, "resources/store/psql/migrations"); err != nil {
		panic(err)
	}
}
