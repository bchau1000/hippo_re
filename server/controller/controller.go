package controller

import (
	"fmt"
	log "hippo/logging"
	middle "hippo/middleware"
	"hippo/service"
	"net/http"
)

type Controller struct {
	VersionController VersionController
	UserController    UserController
}

func (c Controller) HandleFunc(basePath string) {
	log.Info.Printf("Assigning endpoints to controllers")
	urlPathFormat := basePath + "%s"

	http.HandleFunc(
		fmt.Sprintf(urlPathFormat, "version"),
		middle.Wrap(c.VersionController.GetVersion, middle.RequestLogger()))

	http.HandleFunc(
		fmt.Sprintf(urlPathFormat, "user"),
		c.UserController.GetUsers)
}

func NewController(service service.Service) Controller {
	log.Info.Printf("Initializing controller dependencies")
	return Controller{
		VersionController: NewVersionController(service.PingService),
		UserController:    NewUserController(service.UserService),
	}
}
