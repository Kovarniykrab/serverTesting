package service

import (
	"log/slog"

	"github.com/Kovarniykrab/serverTesting/configs"
	"github.com/Kovarniykrab/serverTesting/database"
	"github.com/Kovarniykrab/serverTesting/pkg"
)

type Service struct {
	cfg        *configs.Config
	logger     *slog.Logger
	repositopy *database.Repository
	JWTService *pkg.JWTService
}

func New(cfg *configs.Config, logger *slog.Logger, re *database.Repository) *Service {
	jwtService := pkg.NewJWT(cfg.JWT.SecretKey)

	return &Service{
		cfg:        cfg,
		logger:     logger,
		repositopy: re,
		JWTService: jwtService,
	}
}
