package handler

import (
	"encoding/json"
	"hippo/logging"
	"hippo/model"
	"hippo/service"
	"net/http"
)

type VersionHandler struct {
	PingService service.PingService
}

func (vh *VersionHandler) GetVersion(resp http.ResponseWriter, request *http.Request) {
	vh.PingService.PingDatabase()
	version := &model.Version{
		Version: "1.0",
		Status:  "OK",
	}

	data, err := json.Marshal(version)

	if err != nil {
		logging.Fatal.Fatalf("Fatal error encountered while decoding json: %v", err)
	}

	logging.Info.Print("Retrieving server version...")

	resp.WriteHeader(http.StatusOK)
	resp.Write(data)
}

func NewVersionHandler(pingService service.PingService) VersionHandler {
	return VersionHandler{
		PingService: pingService,
	}
}
