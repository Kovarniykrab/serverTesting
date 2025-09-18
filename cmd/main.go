package main

import (
	"database/sql"
	"embed"
	"log/slog"
	"os"
	"strconv"

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

	slog.SetLogLoggerLevel(conf.Web.LogLevel)

	var _ = handlers.RegisterUserHandler
	slog.Info("API server started",
		"host", conf.Web.Host,
		"port", conf.Web.Port)
	r := routers.GetRouter()

	db, err := database.New(conf.PSQL, slog.Default())
	if err != nil {
		slog.Error("Database connection failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	migrate(conf.PSQL)

	certFile := conf.Web.SSLSertPath
	keyFile := conf.Web.SSLKeyPath

	_, certErr := os.Stat(certFile)
	_, keyErr := os.Stat(keyFile)
	address := conf.Web.Host + ":" + strconv.Itoa(conf.Web.Port)

	if certErr == nil && keyErr == nil {
		slog.Info("SSL found. Starting HTTPS server",
			"address", address,
			"certFile", certFile,
			"keyFile", keyFile)

		err := fasthttp.ListenAndServeTLS(address, certFile, keyFile, r.Handler)
		if err != nil {
			slog.Error("HTTPS server failed", "error", err)
			os.Exit(1)
		}
	} else {
		slog.Info("SSL certificates NOT FOUND. Starting HTTP server",
			"address", address,
			"certErr", certErr,
			"keyErr", keyErr)
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
