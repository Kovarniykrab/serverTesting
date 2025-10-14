package service

import (
	"log/slog"

	"github.com/Kovarniykrab/serverTesting/configs"
	"github.com/Kovarniykrab/serverTesting/database"
)

type Service struct {
	cfg        *configs.Config
	logger     *slog.Logger
	re         *database.Repository
	JWTService *JWTService
}

func New(cfg *configs.Config, logger *slog.Logger, re *database.Repository) *Service {

	return &Service{
		cfg:    cfg,
		logger: logger,
		re:     re,
	}
}
