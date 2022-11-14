package controller

import (
	"fmt"
	key "hippo/common/key"
	"hippo/model"
	"net/http"
)

// Controller to handle response for an error -- no need to give the details, just a guid to attach to the error
func ServerErrorHandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusInternalServerError)
	errorResponse(resp, req, "An internal error occurred while fulfilling the request")
}

func ClientErrorHandler(resp http.ResponseWriter, req *http.Request, message string) {
	resp.WriteHeader(http.StatusBadRequest)
	errorResponse(resp, req, message)
}

func CreatedHandler(resp http.ResponseWriter, req *http.Request, message string) {
	resp.WriteHeader(http.StatusCreated)
	resp.Write(model.Success(message))
}

func ResolveError(resp http.ResponseWriter, req *http.Request, err error) {
}

func errorResponse(resp http.ResponseWriter, req *http.Request, message string) {
	requestId := fmt.Sprint(req.Context().Value(key.RequestId))
	resp.Write(model.Error(message, requestId))
}
