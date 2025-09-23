package handlers

import (
	"log/slog"

	"github.com/Kovarniykrab/serverTesting/application/service"
	"github.com/Kovarniykrab/serverTesting/configs"
	"github.com/valyala/fasthttp"
)

type App struct {
	cfg     *configs.Config
	service *service.App
	logs    *slog.Logger
}

func New(ctx *fasthttp.RequestCtx, cfg *configs.Config, service *service.App, logs *slog.Logger) *App {
	app := &App{cfg: cfg, service: service, logs: logs}

	return app
}
