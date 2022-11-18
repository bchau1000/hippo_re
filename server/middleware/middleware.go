package middleware

import (
	"hippo/logging"
	"hippo/service"
)

type Middleware struct {
	AuthUserMiddle       AuthUserMiddle
	RequestLoggerMiddle  RequestLoggerMiddle
	ResponseHeaderMiddle ResponseHeaderMiddle
}

func NewMiddleware(service service.Service) Middleware {
	logging.Info.Print("Initializing middleware dependencies")
	return Middleware{
		AuthUserMiddle:       NewAuthUserMiddle(service.UserService),
		RequestLoggerMiddle:  NewRequestLoggerMiddle(),
		ResponseHeaderMiddle: NewResponseHeaderMiddle(),
	}
}
