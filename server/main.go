package main

import (
	"fmt"
	"hippo/config"
	"hippo/database"
	"hippo/handler"
	"hippo/logging"
	"hippo/service"
	"net/http"
)

func main() {
	conf := config.GetConfig()

	// Initialize MySQL Driver
	database.Init(conf)

	// Initialize handlers, services, and repositories
	services := service.NewService()
	handlers := handler.NewHandler(services)

	// Initialize endpoints
	handlers.HandleFunc(conf.Server.BasePath)

	// Initialize and start server
	initServer(conf)
}

func initServer(config *config.Config) {
	address := fmt.Sprintf(":%d", config.Server.Port)

	logging.Info.Printf("Starting server on port: %d", config.Server.Port)
	if err := http.ListenAndServe(address, nil); err != nil {
		logging.Fatal.Fatalf("Fatal error encountered while starting server: %v", err)
	}
}
