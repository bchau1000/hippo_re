package main

import (
	"fmt"
	"hippo/config"
	"hippo/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(f)

	if err != nil {
		log.Fatalf("Fatal error opening server log: %v", err)
	}
	defer f.Close()

	conf := config.GetConfig()
	basePath := conf.Server.BasePath + "%s"

	handler.Init(basePath)

	address := fmt.Sprintf(":%d", conf.Server.Port)
	log.Printf("Server listening on: %s", address)
	err = http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf("Fatal error listening and serving: %v", err)
	}
}
