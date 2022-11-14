package controller

import (
	"encoding/json"
	"hippo/common/errormsg"
	"hippo/logging"
	"hippo/model"
	"hippo/service"
	"net/http"
)

type userResponse struct {
	Users []model.User `json:"users"`
}

type registerUserRequest struct {
	User model.UserToCreate `json:"user"`
}

type UserController struct {
	UserService service.UserService
}

func (uc *UserController) GetUsers(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	users, err := uc.UserService.GetByIds(ctx)
	if err != nil {
		ErrorHandler(resp, req, err)
		return
	}

	user, err := uc.UserService.GetByEmail(ctx, "admin@email.com")
	if err != nil {
		ErrorHandler(resp, req, err)
		return
	}

	users = append(users, user)
	userReponse := userResponse{
		Users: users,
	}

	data, err := json.Marshal(userReponse)
	if err != nil {
		logging.Error.Print(errormsg.FormatError(ctx, errormsg.ConvertJson, err))
		ErrorHandler(resp, req, err)
		return
	}

	resp.Write(data)
}

func (uc *UserController) RegisterUser(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	registerRequest := &registerUserRequest{}

	err := json.NewDecoder(req.Body).Decode(registerRequest)
	if err != nil {
		logging.Error.Printf(errormsg.FormatError(ctx, errormsg.DecodeJson, err))
		ErrorHandler(resp, req, err)
		return
	}

	err = uc.UserService.RegisterUser(ctx, registerRequest.User)
	if err != nil {
		ErrorHandler(resp, req, err)
		return
	}

	StatusCreatedHandler(resp, req, "User successfully registered!")
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}
