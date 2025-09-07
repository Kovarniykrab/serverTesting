package main

import (
	"fmt"

	"github.com/Kovarniykrab/serverTesting/api/handlers"
	"github.com/Kovarniykrab/serverTesting/api/routers"
	_ "github.com/Kovarniykrab/serverTesting/docs"
	"github.com/valyala/fasthttp"
)

// @title          TestUser API
// @version        0.5
// @description    API для управления пользователями
// @host           localhost:8080
// @BasePath       /api
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

func main() {
	var _ = handlers.RegisterUserHandler
	fmt.Println("API server started on :8080")
	r := routers.GetRouter()
	fasthttp.ListenAndServe(":8080", r.Handler)

}
