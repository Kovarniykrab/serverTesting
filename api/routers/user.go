package routers

import "github.com/fasthttp/router"

func (app *App) userRoutes(user *router.Group) {
	user.GET("/profile/{id}", app.GetUserHandler)
	user.POST("/register", app.RegisterUserHandler)
	user.PUT("/change_password/{id}", app.AuthMiddleware(app.UpdatePasswordHandler))
	user.PUT("/change_user/{id}", app.ChangeUserHandler)
	user.DELETE("/api/user/delete/{id}", app.AuthMiddleware(app.DeleteUserHandler))
}
