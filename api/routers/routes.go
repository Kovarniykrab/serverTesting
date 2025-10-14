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
	user.GET("/profile/{id}", app.GetUserHandler)
	user.POST("/register", app.RegisterUserHandler)
	user.PUT("/change_password/{id}", app.AuthMiddleware(app.UpdatePasswordHandler))
	user.PUT("/change_user/{id}", app.ChangeUserHandler)
	user.DELETE("/api/user/delete/{id}", app.AuthMiddleware(app.DeleteUserHandler))
	user.POST("/logout", app.AuthMiddleware(app.LogoutUserHandler))
	user.POST("/login", app.AuthUserHandler)
	user.GET("/api/user/check", app.AuthMiddleware(app.CheckHandler))

	return routers

}
