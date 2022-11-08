package main

import (
	"fmt"
	"hippo/config"
	"hippo/handler"
	"hippo/logging"
	"log"
	"net/http"
)

func main() {
	log.SetOutput(logging.Logfile)
	defer logging.Logfile.Close()

	conf := config.GetConfig()
	basePath := conf.Server.BasePath + "%s"

	handler.Init(basePath)

	address := fmt.Sprintf(":%d", conf.Server.Port)
	logging.Info.Printf("Server listening on port: %d", conf.Server.Port)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf("Fatal error listening and serving: %v", err)
	}
}
