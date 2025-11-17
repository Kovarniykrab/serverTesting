package routers

import (
	"context"
	"log/slog"

	"github.com/Kovarniykrab/serverTesting/api/handlers"
	"github.com/Kovarniykrab/serverTesting/configs"
	"github.com/Kovarniykrab/serverTesting/database"
	"github.com/Kovarniykrab/serverTesting/internal/service"
	"github.com/fasthttp/router"
	swagger "github.com/swaggo/fasthttp-swagger"
	"github.com/valyala/fasthttp"
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
func (app *App) applyMiddlewares(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	handler = app.PanicRecoveryMiddleware(handler)
	handler = app.LoggingMiddleware(handler)
	handler = app.CorsMiddleware(handler)

	return handler
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
func (app *App) GetHandler() fasthttp.RequestHandler {
	r := app.GetRouter()
	return app.applyMiddlewares(r.Handler)
}

// мидлвары на верхнем уровне
// разделить роуты на разные файлы, где милдвар нужен и не нужен
// логаут положить в routes.go, мидлвар ему не нужен
// корсы мидлвар подобные маяку
// обработка мидлвар пиники, неправильный метод(такой метод недоступен) и запрос.
// сделать уровни мидл вар как в маяк /app/router/app.go (ЛЕГАСИ)
