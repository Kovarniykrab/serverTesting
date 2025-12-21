package routers

import "github.com/fasthttp/router"

func (app *App) userRoutes(user *router.Group) {
	user.POST("/logout", app.AuthMiddleware(app.Logout))
	user.GET("/check", app.AuthMiddleware(app.Check))
	user.PUT("/change_password/{id}", app.AuthMiddleware(app.UpdatePassword))
	user.PUT("/change_user/{id}", app.AuthMiddleware(app.Change))
	user.DELETE("/delete/{id}", app.AuthMiddleware(app.Delete))
}
