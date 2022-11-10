package main

import (
	"fmt"
	"hippo/config"
	"hippo/controller"
	"hippo/database"
	"hippo/logging"
	"hippo/repository"
	"hippo/service"
	"net/http"
)

func main() {
	// Read in configurations
	conf := config.GetConfig()

	// Initialize MySQL Driver
	database.Init(conf)

	// Initialize repositories, services, and controllers
	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	controllers := controller.NewController(services)

	// Initialize endpoints
	controllers.HandleFunc(conf.Server.BasePath)

	// Start the server
	startServer(conf)
}

func startServer(config *config.Config) {
	address := fmt.Sprintf(":%d", config.Server.Port)

	logging.Info.Printf("Starting server on port: %d", config.Server.Port)
	if err := http.ListenAndServe(address, nil); err != nil {
		logging.Fatal.Fatalf("Fatal error encountered while starting server: %v", err)
	}
}
