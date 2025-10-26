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
	ctx := context.Background()

	conf := configs.Config{}
	parser := flags.NewParser(&conf, flags.Default)
	if _, err := parser.Parse(); err != nil {
		panic(err)
	}

	log := initLogger(conf)
	log.Info("Starting app")

	log.Info("starting migrate")
	migrate(conf.PSQL)

	r := routers.New(ctx, &conf, log)
	log.Info("routers init")

	address := conf.Web.Host + ":" + strconv.Itoa(conf.Web.Port)

	server := &fasthttp.Server{
		Handler:     r.GetRouter().Handler,
		ReadTimeout: 10 * time.Second,
	}

	if conf.Web.SSLSertPath != "" && conf.Web.SSLKeyPath != "" {
		log.Info("SSL found. Starting HTTPS server",
			"address", address,
			"certFile", conf.Web.SSLSertPath,
			"keyFile", conf.Web.SSLKeyPath)

		err := server.ListenAndServeTLS(address, conf.Web.SSLSertPath, conf.Web.SSLKeyPath)
		if err != nil {
			log.Error("HTTPS server failed", "error", err)
			os.Exit(1)
		}
	} else {
		log.Info("SSL certificates NOT FOUND. Starting HTTP server",
			"address", address,
			"certFile", conf.Web.SSLSertPath,
			"keyFile", conf.Web.SSLKeyPath)
		err := server.ListenAndServe(address)
		if err != nil {
			log.Error("HTTP server failed: %v", "error", err)
			os.Exit(1)
		}
		log.Info("Server starting on:")
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

var embedMigrations embed.FS //nolint:unused

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

// в сервисе проверки регистрации поднять в хендлеры +
// разделить ошибки и не делать их вместе и не дублировать проверки ?
//sql no found маяк посмотреть ---
// в апдейт проверить пользователя по id, если есть, то только тогда апдейт +
// getUserById перевести к единому стилю +
// в регистрации датабейз убрать форму и перенести ее в сервис +
// доставать сертификаты через енвы +-
// проверка сертификатов только на нил и все +
// поправить свагер и хендлеры +
//доменные модели добавить required +
//
