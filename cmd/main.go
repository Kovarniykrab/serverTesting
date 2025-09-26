package main

import (
	"context"
	"database/sql"
	"embed"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	embedServer "github.com/Kovarniykrab/serverTesting"
	"github.com/Kovarniykrab/serverTesting/api/routers"
	"github.com/Kovarniykrab/serverTesting/configs"
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

	log := initLogger(conf)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := routers.New(ctx, &conf, log)

	migrate(conf.PSQL)

	certFile := conf.Web.SSLSertPath
	keyFile := conf.Web.SSLKeyPath

	_, certErr := os.Stat(certFile)
	_, keyErr := os.Stat(keyFile)
	address := conf.Web.Host + ":" + strconv.Itoa(conf.Web.Port)

	server := &fasthttp.Server{
		Handler:     r.GetRouter().Handler,
		ReadTimeout: 10 * time.Second,
	}

	if certErr == nil && keyErr == nil {
		slog.Info("SSL found. Starting HTTPS server",
			"address", address,
			"certFile", certFile,
			"keyFile", keyFile)

		err := server.ListenAndServeTLS(address, certFile, keyFile)
		if err != nil {
			slog.Error("HTTPS server failed", "error", err)
			os.Exit(1)
		}
	} else {
		slog.Info("SSL certificates NOT FOUND. Starting HTTP server",
			"address", address,
			"certErr", certErr,
			"keyErr", keyErr)
		err := server.ListenAndServe(address)
		if err != nil {
			slog.Error("HTTP server failed: %v", "error", err)
			os.Exit(1)
		}
		slog.Info("Server starting on:")
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Info("Shutting down")

	if err := server.Shutdown(); err != nil {
		log.Error("Shutting down", "error", err)
	}
	log.Info("server stoped")

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

func initLogger(conf configs.Config) *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: conf.Web.LogLevel,
	}))

	return logger
}

//контекст нормальные
//запустить приложение с хендлерами
