package controller

import (
	"encoding/json"
	"hippo/common/errormsg"
	"hippo/logging"
	"hippo/model"
	"hippo/service"
	"net/http"
)

type UserController struct {
	userService service.UserService
}

type authUserRequest struct {
	IdToken string `json:"idToken"`
}

type registerUserRequest struct {
	Request model.Request
	User    model.UserToCreate `json:"user"`
}

func (uc *UserController) AuthUser(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	cookie, err := req.Cookie("idToken")
	if err != nil {
		ErrorHandler(resp, req, err)
		return
	}

	idToken := cookie.Value

	_, err = uc.userService.AuthUser(ctx, idToken)
	if err != nil {
		ErrorHandler(resp, req, err)
		return
	}

	resp.Write([]byte("Success! User is logged in"))
}

func (uc *UserController) LoginUser(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	authUser := &authUserRequest{}
	json.
		NewDecoder(req.Body).
		Decode(authUser)

	user, err := uc.userService.AuthUser(ctx, authUser.IdToken)

	if err != nil {
		ErrorHandler(resp, req, err)
		return
	}

	http.SetCookie(
		resp,
		&http.Cookie{
			Name:     "idToken",
			Value:    authUser.IdToken,
			MaxAge:   0,
			Path:     "/",
			HttpOnly: true})

	data, err := json.Marshal(user)
	if err != nil {
		logging.Error.Printf(
			errormsg.FormatError(
				ctx,
				errormsg.DecodeJson,
				err))
		ErrorHandler(resp, req, err)
		return
	}

	resp.Write(data)
}

func (uc *UserController) RegisterUser(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	registerRequest := &registerUserRequest{}

	err := json.
		NewDecoder(req.Body).
		Decode(registerRequest)
	if err != nil {
		logging.Error.Printf(
			errormsg.FormatError(
				ctx,
				errormsg.DecodeJson,
				err))
		ErrorHandler(resp, req, err)
		return
	}

	err = uc.userService.RegisterUser(ctx, registerRequest.User)
	if err != nil {
		ErrorHandler(resp, req, err)
		return
	}

	StatusCreatedHandler(resp, req, "User successfully registered!")
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		userService: userService,
	}
}
