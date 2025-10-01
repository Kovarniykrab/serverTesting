package service

import (
	"context"
	"log/slog"

	"github.com/Kovarniykrab/serverTesting/configs"
	"github.com/Kovarniykrab/serverTesting/database"
)

type Service struct {
	cfg    *configs.Config
	logger *slog.Logger
	ctx    context.Context
	re     *database.Repository
}

func New(ctx context.Context, cfg *configs.Config, logger *slog.Logger, re *database.Repository) *Service {

	return &Service{
		cfg:    cfg,
		logger: logger,
		ctx:    ctx,
		re:     re,
	}
}
