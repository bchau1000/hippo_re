package middleware

import (
	"context"
	key "hippo/common/key"
	"hippo/logging"
	"net/http"

	"github.com/google/uuid"
)

type RequestLoggerMiddle struct {
}

// Logs the request path and ID
func (rl *RequestLoggerMiddle) Wrap() MiddlewareWrapper {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(resp http.ResponseWriter, req *http.Request) {
			id, _ := uuid.NewRandom()
			logging.Info.Printf("Handling request - %s: %s", id, req.URL.Path)
			req = req.WithContext(
				context.WithValue(
					req.Context(),
					key.RequestId, id))

			next.ServeHTTP(resp, req)
		}
	}
}

func NewRequestLoggerMiddle() RequestLoggerMiddle {
	return RequestLoggerMiddle{}
}
