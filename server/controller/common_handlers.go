package controller

import (
	"fmt"
	"hippo/common/errormsg"
	key "hippo/common/key"
	"hippo/model"
	"net/http"
)

func StatusCreatedHandler(resp http.ResponseWriter, req *http.Request, message string) {
	resp.WriteHeader(http.StatusCreated)
	resp.Write(model.Success(message))
}

func ErrorHandler(resp http.ResponseWriter, req *http.Request, err error) {
	errorMsg := errormsg.IsClientError(err)
	if errorMsg != "" {
		clientErrorHandler(resp, req, errorMsg)
		return
	}
	serverErrorHandler(resp, req)
}

// Controller to handle response for an error -- no need to give the details, just a guid to attach to the error
func serverErrorHandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusInternalServerError)
	errorResponse(resp, req, "An internal error occurred while fulfilling the request")
}

func clientErrorHandler(resp http.ResponseWriter, req *http.Request, message string) {
	resp.WriteHeader(http.StatusBadRequest)
	errorResponse(resp, req, message)
}

func errorResponse(resp http.ResponseWriter, req *http.Request, message string) {
	requestId := fmt.Sprint(req.Context().Value(key.RequestId))
	resp.Write(model.Error(message, requestId))
}
