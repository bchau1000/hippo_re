package handler

import (
	"encoding/json"
	"hippo/logging"
	"hippo/model"
	"net/http"
)

func GetVersion(resp http.ResponseWriter, request *http.Request) {
	version := &model.Version{
		Version: "1.0",
		Status:  "OK",
	}

	data, err := json.Marshal(version)

	if err != nil {
		logging.Error.Print(err)
	}

	logging.Info.Print("Retrieving server version...")

	resp.WriteHeader(http.StatusOK)
	resp.Write(data)
}
