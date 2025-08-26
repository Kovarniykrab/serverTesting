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
// @host           localhost:8080
// @BasePath       /api
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	apiStop := make(chan struct{})
	swaggerStop := make(chan struct{})

	go apiServ(apiStop)

	go swag(swaggerStop)

	<-done
	log.Println("Server's shut down")
}

func apiServ(stop <-chan struct{}) {
	log.Println("API server starting on :8080")
	server := fasthttp.Server{
		Handler: routers.GetRouter().Handler,
	}

	go func() {
		if err := server.ListenAndServe(":8080"); err != nil {
			log.Printf("API server failed: %v", err)
		}
	}()

	<-stop
	server.Shutdown()
}

func swag(stop <-chan struct{}) {
	log.Println("Swagger server starting on :8085")
	log.Println("http://localhost:8085/swagger/index.html")
	swaggerRouter := router.New()
	swaggerRouter.GET("/swagger/{filepath:*}", swagger.WrapHandler())

	server := &fasthttp.Server{
		Handler: swaggerRouter.Handler,
	}

	go func() {
		if err := server.ListenAndServe(":8085"); err != nil {
			log.Printf("Swagger server failed: %v", err)
		}
	}()

	<-stop
	server.Shutdown()
}

//залогировать с помощью пакета log
//запустить каждый порт в отдельной гарутине
//че не так блять
