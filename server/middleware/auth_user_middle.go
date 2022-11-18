package middleware

import (
	"hippo/service"
	"net/http"
)

type AuthUserMiddle struct {
	userService service.UserService
}

func (au *AuthUserMiddle) Wrap() MiddlewareWrapper {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(resp http.ResponseWriter, req *http.Request) {

		}
	}
}

func NewAuthUserMiddle(userService service.UserService) AuthUserMiddle {
	return AuthUserMiddle{
		userService: userService,
	}
}
