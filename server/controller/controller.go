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

	commonMiddleware := []middle.Middleware{middle.RequestLogger()}

	http.HandleFunc(
		fmt.Sprintf(urlPathFormat, "version"),
		middle.Wrap(c.VersionController.GetVersion, commonMiddleware...))

	http.HandleFunc(
		fmt.Sprintf(urlPathFormat, "user"),
		middle.Wrap(c.UserController.GetUsers, commonMiddleware...))
}

func NewController(service service.Service) Controller {
	log.Info.Printf("Initializing controller dependencies")
	return Controller{
		VersionController: NewVersionController(service.PingService),
		UserController:    NewUserController(service.UserService),
	}
}
