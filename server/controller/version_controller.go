package controller

import (
	"encoding/json"
	log "hippo/logging"
	"hippo/model"
	"hippo/service"
	"net/http"
)

type VersionController struct {
	PingService service.PingService
}

func (vh *VersionController) GetVersion(resp http.ResponseWriter, req *http.Request) {
	pingSuccess := vh.PingService.PingDatabase(req.Context())

	status := "OK"
	if !pingSuccess {
		status = "BadRequest"
		resp.WriteHeader(http.StatusBadRequest)
	}

	version := &model.Version{
		Version: "1.0",
		Status:  status,
	}

	data, err := json.Marshal(version)

	if err != nil {
		log.Fatal.Fatalf("Fatal error encountered while decoding json: %v", err)
	}

	resp.Write(data)
}

func NewVersionController(pingService service.PingService) VersionController {
	return VersionController{
		PingService: pingService,
	}
}
