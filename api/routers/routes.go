package routers

import (
	"context"
	"log/slog"

	"github.com/Kovarniykrab/serverTesting/api/handlers"
	"github.com/Kovarniykrab/serverTesting/application/service"
	"github.com/Kovarniykrab/serverTesting/configs"
	"github.com/Kovarniykrab/serverTesting/database"
	"github.com/fasthttp/router"
	swagger "github.com/swaggo/fasthttp-swagger"
)

type App struct {
	handlers.App
	config *configs.Config
	logger *slog.Logger
}

func New(ctx context.Context, cfg *configs.Config, log *slog.Logger) *App {

	db, err := database.New(ctx, cfg.PSQL, log)
	if err != nil {
		panic(err)
	}

	srv := service.New(cfg, log, db)
	hand := handlers.New(cfg, srv, log)

	return &App{
		App:    *hand,
		config: cfg,
		logger: log,
	}
}

func (app *App) GetRouter() *router.Router {
	routers := router.New()
	api := routers.Group("/api")
	swaggerRouter := routers.Group("/info")
	swaggerRouter.GET("/swagger/{any:*}", swagger.WrapHandler())

	user := api.Group("/user")
	app.userRoutes(user)
	app.authRoutes(user)

	return routers

}
