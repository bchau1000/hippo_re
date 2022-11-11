package middleware

import (
	"context"
	"hippo/logging"
	"hippo/model"
	"net/http"

	"github.com/google/uuid"
)

// Logs the request path and ID
func RequestLogger() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(resp http.ResponseWriter, req *http.Request) {
			id, _ := uuid.NewRandom()
			logging.Info.Printf("Handling request - %s: %s", id, req.URL.Path)
			req = req.WithContext(
				context.WithValue(
					req.Context(),
					model.RequestId{Id: "requestId"}, id))

			next.ServeHTTP(resp, req)
		}
	}
}
