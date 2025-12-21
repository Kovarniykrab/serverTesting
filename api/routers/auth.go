package routers

import "github.com/fasthttp/router"

func (app *App) authRoutes(user *router.Group) {
	user.POST("/login", app.Auth)
	user.GET("/profile/{id}", app.Get)
	user.POST("/register", app.Register)

}
