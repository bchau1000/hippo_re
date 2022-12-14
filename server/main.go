package main

import (
	"context"
	"fmt"
	"hippo/config"
	"hippo/controller"
	"hippo/database"
	"hippo/logging"
	"hippo/middleware"
	"hippo/repository"
	"hippo/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Read in configurations, set context
	conf := config.GetConfig()
	ctx := context.Background()

	// Initialize database connection
	db := database.NewDatabase(ctx, conf)

	// Initialize Firebase Auth client
	fbAuth := database.NewFirebaseAuth(ctx, conf)

	// Initialize router
	router := mux.NewRouter()

	// Initialize repositories, services, and controllers
	repositories := repository.NewRepository(db, fbAuth)
	services := service.NewService(repositories)
	middleware := middleware.NewMiddleware(services)
	controllers := controller.NewController(services, middleware)

	// Delegate endpoints to controller methods
	controllers.HandleFunc(conf.Server.BasePath, router)

	// Start the server
	startServer(conf, router)
}

func startServer(config *config.Config, router *mux.Router) {
	address := fmt.Sprintf(":%d", config.Server.Port)

	logging.Info.Printf("Starting server on port: %d", config.Server.Port)
	if err := http.ListenAndServe(address, router); err != nil {
		logging.Fatal.Fatalf("Fatal error encountered while starting server: %v", err)
	}
}
