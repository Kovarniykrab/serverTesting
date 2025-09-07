package routers

import (
	"github.com/Kovarniykrab/serverTesting/api/handlers"
	"github.com/fasthttp/router"
	swagger "github.com/swaggo/fasthttp-swagger"
)

func GetRouter() *router.Router {
	routers := router.New()
	api := routers.Group("/api")
	swaggerRouter := routers.Group("/swagger/")
	swaggerRouter.GET("/swagger/", swagger.WrapHandler())

	user := api.Group("/user")
	user.GET("/profile/{id}", handlers.GetUserHandler)
	user.POST("/register", handlers.RegisterUserHandler)
	user.PUT("/changePassword{id}", handlers.UpdatePasswordHandler)
	user.PUT("/changeUser", handlers.ChangeUserHandler)
	user.DELETE("/delete/{id}", handlers.DeleteUserHandler)
	user.POST("/logout", handlers.LogoutUserHandler)
	user.POST("/login", handlers.AuthUserHandler)

	return routers

}
