package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Kovarniykrab/serverTesting/api/handlers"
	"github.com/Kovarniykrab/serverTesting/api/routers"
	_ "github.com/Kovarniykrab/serverTesting/docs"
	"github.com/valyala/fasthttp"
)

// @title          TestUser API
// @version        0.5
// @description    API для управления пользователями
// @host           wednode.ru:8080
// @BasePath       /
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

func main() {
	var _ = handlers.RegisterUserHandler
	fmt.Println("API server started on :8080")
	r := routers.GetRouter()

	certDirectory := "/etc/letsencrypt/live/wednode.ru"
	certFile := filepath.Join(certDirectory, "fullchain.pem")
	keyFile := filepath.Join(certDirectory, "privkey.pem")

	_, certErr := os.Stat(certFile)
	_, keyErr := os.Stat(keyFile)

	if certErr == nil && keyErr == nil {
		fmt.Printf("SSL found. Starting HTTPS server on :8080")
		err := fasthttp.ListenAndServeTLS(":8080", certFile, keyFile, r.Handler)
		if err != nil {
			log.Fatal("HTTPS server failed", err)
		}
	} else {
		fmt.Printf("SSL certificates NOT FOUND. Starting HTTP server on :8080\n")
		err := fasthttp.ListenAndServe(":8080", r.Handler)
		if err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}

}
