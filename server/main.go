package main

import (
	"fmt"
	"hippo/config"
	"hippo/database"
	"hippo/handler"
	"hippo/logging"
	"net/http"
)

func main() {
	conf := config.GetConfig()
	basePath := conf.Server.BasePath + "%s"

	// Initialize MySQL Driver
	database.Init(conf)

	// Initialize all handlers and endpoints
	handler.Init(basePath)

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
