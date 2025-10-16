package routers

import "github.com/fasthttp/router"

func (app *App) authRoutes(user *router.Group) {
	user.POST("/logout", app.AuthMiddleware(app.LogoutUserHandler))
	user.POST("/login", app.AuthUserHandler)
	user.GET("/check", app.AuthMiddleware(app.CheckHandler))
}
