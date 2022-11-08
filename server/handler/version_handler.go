package handler

import (
	"encoding/json"
	"fmt"
	"hippo/logging"
	"hippo/model"
	"net/http"
)

type VersionHandler struct{}

func (vh *VersionHandler) GetVersion(resp http.ResponseWriter, request *http.Request) {
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

func NewVersionHandler(hippoBasePath string) {
	vh := &VersionHandler{}
	http.HandleFunc(
		fmt.Sprintf(hippoBasePath, "version"),
		vh.GetVersion)
}
