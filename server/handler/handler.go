package handler

import (
	"fmt"
	"hippo/service"
	"net/http"
)

type Handler struct {
	VersionHandler VersionHandler
}

func (h Handler) HandleFunc(basePath string) {
	urlPathFormat := basePath + "%s"

	http.HandleFunc(
		fmt.Sprintf(urlPathFormat, "version"),
		h.VersionHandler.GetVersion)
}

func NewHandler(service service.Service) Handler {
	return Handler{
		VersionHandler: NewVersionHandler(service.PingService),
	}
}
