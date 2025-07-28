package cmd

import (
	"github.com/Kovarniykrab/serverTesting/myApp/api/handlers"
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
}
