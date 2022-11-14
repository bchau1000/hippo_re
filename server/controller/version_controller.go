package controller

import (
	"encoding/json"
	"hippo/common"
	"hippo/logging"
	"hippo/model"
	"hippo/service"
	"net/http"
)

type VersionController struct {
	PingService service.PingService
}

func (vh *VersionController) GetVersion(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	pingSuccess, err := vh.PingService.PingDatabase(ctx)

	if !pingSuccess || err != nil {
		ServerErrorHandler(resp, req)
		return
	}

	version := &model.Version{
		Version: "1.0",
		Status:  "OK",
	}

	data, err := json.Marshal(version)

	if err != nil {
		logging.Error.Printf(common.FormatError(ctx, common.ServerError.ConvertJson, err))
		ServerErrorHandler(resp, req)
		return
	}

	resp.Write(data)
}

func NewVersionController(pingService service.PingService) VersionController {
	return VersionController{
		PingService: pingService,
	}
}
