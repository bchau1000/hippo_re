package controller

import (
	"net/http"
)

type ErrorController struct {
}

// Controller to handle response for an error -- no need to give the details, just a guid to attach to the error
func (ec *ErrorController) ServerErrorHandler(resp http.ResponseWriter, req *http.Request) {

}

// Any errors should propagate down to the controller level
func NewErrorController() ErrorController {
	return ErrorController{}
}
