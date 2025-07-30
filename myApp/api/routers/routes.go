package routers

import (
	"github.com/Kovarniykrab/serverTesting/myApp/api/handlers"
	"github.com/fasthttp/router"
)

func GetRouter() *router.Router {
	routers := router.New()
	api := routers.Group("/api")

	user := api.Group("/user")
	user.GET("/profile/{id}", handlers.GetUserHandler)
	user.POST("/register", handlers.RegisterUserHandler)
	user.PUT("/update", handlers.UpdateUserHandler)
	user.DELETE("/delete{id}", handlers.DeleteUserHandler)
	user.POST("/logout", handlers.LogoutUserHandler)
	user.POST("/login", handlers.AuthUserHandler)

	return routers
}
