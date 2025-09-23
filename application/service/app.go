package service

import (
	"context"
	"log/slog"

	"github.com/Kovarniykrab/serverTesting/configs"
	"github.com/Kovarniykrab/serverTesting/domain"
)

type App struct {
	cfg      *configs.Config
	logger   *slog.Logger
	Database Database
	ctx      context.Context
}

type Database interface {
	RegisterUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id int) error
	ChangeUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, id int) (*domain.User, error)
}

func New(ctx context.Context, cfg *configs.Config, logger *slog.Logger, Database Database) *App {

	return &App{
		cfg:      cfg,
		logger:   logger,
		Database: Database,
		ctx:      ctx,
	}
}
