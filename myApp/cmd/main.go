package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go apiServ()

	go swag()
	<-done
	log.Println("Server's shut down")
}

func apiServ() {
	log.Println("API server starting on :8080")
	r := routers.GetRouter()
	if err := fasthttp.ListenAndServe(":8080", r.Handler); err != nil {
		log.Fatalf("API server failed: %v", err)
	}
}

func swag() {
	log.Println("Swagger server starting on :8085")
	log.Println("")
	swaggerRouter := router.New()
	swaggerRouter.GET("/swagger/{filepath:*}", swagger.WrapHandler())
	if err := fasthttp.ListenAndServe(":8085", swaggerRouter.Handler); err != nil {
		log.Fatalf("Swagger server failed: %v", err)
	}
}

//залогировать с помощью пакета log
//запустить каждый порт в отдельной гарутине
