package controller

import (
	"fmt"
	log "hippo/logging"
	"hippo/service"
	"net/http"
)

type Controller struct {
	VersionController VersionController
}

func (c Controller) HandleFunc(basePath string) {
	log.Info.Printf("Assigning endpoints to controllers")
	urlPathFormat := basePath + "%s"

	http.HandleFunc(
		fmt.Sprintf(urlPathFormat, "version"),
		c.VersionController.GetVersion)
}

func NewController(service service.Service) Controller {
	log.Info.Printf("Initializing controller dependencies")
	return Controller{
		VersionController: NewVersionController(service.PingService),
	}
}
