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
	authUser := &authUserRequest{}
	json.
		NewDecoder(req.Body).
		Decode(authUser)

	user, err := uc.userService.AuthUser(ctx, authUser.IdToken)
	if err != nil {
		ErrorHandler(resp, req, err)
		return
	}

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
