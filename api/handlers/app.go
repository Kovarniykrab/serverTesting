package handlers

import (
	"context"
	"log/slog"

	"github.com/Kovarniykrab/serverTesting/application/service"
	"github.com/Kovarniykrab/serverTesting/configs"
)

type App struct {
	cfg     *configs.Config
	Service *service.Service
	logs    *slog.Logger
}

func New(ctx context.Context, cfg *configs.Config, service *service.Service, logs *slog.Logger) *App {
	app := &App{
		cfg:     cfg,
		Service: service,
		logs:    logs}

	return app
}
