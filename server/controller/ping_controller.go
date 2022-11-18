package controller

import (
	"encoding/json"
	"hippo/common/errormsg"
	"hippo/logging"
	"hippo/service"
	"net/http"
)

type PingController struct {
	pingService service.PingService
}

type Pong struct {
	Version string `json:"version"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (pc *PingController) Ping(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	pingSuccess, err := pc.pingService.PingDatabase(ctx)

	if !pingSuccess || err != nil {
		ErrorHandler(resp, req, err)
		return
	}

	version := &Pong{
		Version: "1.0",
		Status:  "OK",
		Message: "Pong",
	}

	data, err := json.Marshal(version)

	if err != nil {
		logging.Error.Printf(
			errormsg.FormatError(
				ctx,
				errormsg.ConvertJson,
				err))
		ErrorHandler(resp, req, err)
		return
	}

	resp.Write(data)
}

func NewPingController(pingService service.PingService) PingController {
	return PingController{
		pingService: pingService,
	}
}
