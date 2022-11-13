package controller

import (
	"fmt"
	log "hippo/logging"
	middle "hippo/middleware"
	"hippo/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	VersionController VersionController
	UserController    UserController
}

func (c Controller) HandleFunc(basePath string, router *mux.Router) {
	log.Info.Printf("Assigning endpoints to controllers")
	urlPathFormat := basePath + "%s"

	commonMiddleware := []middle.Middleware{middle.RequestLogger(), middle.ResponseHeader()}

	router.HandleFunc(
		fmt.Sprintf(urlPathFormat, "version"),
		middle.Wrap(c.VersionController.GetVersion, commonMiddleware...),
	).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc(
		fmt.Sprintf(urlPathFormat, "user"),
		middle.Wrap(c.UserController.GetUsers, commonMiddleware...),
	).Methods(http.MethodGet, http.MethodOptions)
}

func NewController(service service.Service) Controller {
	log.Info.Printf("Initializing controller dependencies")
	return Controller{
		VersionController: NewVersionController(service.PingService),
		UserController:    NewUserController(service.UserService),
	}
}
