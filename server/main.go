package main

import (
	"fmt"
	"hippo/config"
	"hippo/handler"
	"log"
	"net/http"
)

func main() {
	conf := config.GetConfig()
	basePath := conf.Server.BasePath + "%s"

	handler.Init(basePath)

	address := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}
