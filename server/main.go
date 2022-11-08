package main

import (
	"fmt"
	"hippo/config"
	"hippo/handler"
	"hippo/logging"
	"net/http"
)

func main() {
	conf := config.GetConfig()
	basePath := conf.Server.BasePath + "%s"

	handler.Init(basePath)

	address := fmt.Sprintf(":%d", conf.Server.Port)
	logging.Info.Printf("Server listening on port: %d", conf.Server.Port)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		logging.Fatal.Fatalf("Fatal error listening and serving: %v", err)
	}
}
