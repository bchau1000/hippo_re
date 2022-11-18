package controller

import (
	"fmt"
	log "hippo/logging"
	mw "hippo/middleware"
	"hippo/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	middleware     mw.Middleware
	pingController PingController
	userController UserController
}

func (c Controller) HandleFunc(basePath string, router *mux.Router) {
	log.Info.Printf("Assigning endpoints to controllers")
	urlPathFormat := basePath + "%s"

	commonMiddleware := []mw.IMiddlewareWrapper{
		&c.middleware.RequestLoggerMiddle,
		&c.middleware.ResponseHeaderMiddle,
	}

	router.
		HandleFunc(
			fmt.Sprintf(urlPathFormat, "/ping"),
			mw.Wrap(c.pingController.Ping, commonMiddleware...)).
		Methods(http.MethodGet, http.MethodOptions)

	router.
		HandleFunc(
			fmt.Sprintf(urlPathFormat, "/user/auth"),
			mw.Wrap(c.userController.AuthUser, commonMiddleware...)).
		Methods(http.MethodPost, http.MethodOptions)

	router.
		HandleFunc(
			fmt.Sprintf(urlPathFormat, "/user/register"),
			mw.Wrap(c.userController.RegisterUser, commonMiddleware...)).
		Methods(http.MethodPut, http.MethodOptions)
}

func NewController(service service.Service, middleware mw.Middleware) Controller {
	log.Info.Printf("Initializing controller dependencies")
	return Controller{
		middleware:     middleware,
		pingController: NewPingController(service.PingService),
		userController: NewUserController(service.UserService),
	}
}
