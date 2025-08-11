package main

import (
	"fmt"

	"github.com/Kovarniykrab/serverTesting/api/handlers"
	"github.com/Kovarniykrab/serverTesting/api/routers"
	_ "github.com/Kovarniykrab/serverTesting/docs"
	"github.com/fasthttp/router"
	swagger "github.com/swaggo/fasthttp-swagger"
	"github.com/valyala/fasthttp"
)

// @title          TestUser API
// @version        0.5
// @description    API для управления пользователями
// @host           localhost:8085
// @BasePath       /swagger/
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

func main() {
	var _ = handlers.RegisterUserHandler
	fmt.Println("API server started on :8080")
	go func() {
		r := routers.GetRouter()
		fasthttp.ListenAndServe(":8080", r.Handler)
	}()

	fmt.Println("Swagger server started on :8085")
	swaggerRouter := router.New()

	swaggerRouter.GET("/swagger/*any", swagger.WrapHandler())
	fasthttp.ListenAndServe(":8085", swaggerRouter.Handler)
}
