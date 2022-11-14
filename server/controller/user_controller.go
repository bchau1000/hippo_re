package controller

import (
	"encoding/json"
	"hippo/common"
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
		ServerErrorHandler(resp, req)
		return
	}

	user, err := uc.UserService.GetByEmail(ctx, "admin@email.com")
	if err != nil {
		ServerErrorHandler(resp, req)
		return
	}

	users = append(users, user)
	userReponse := userResponse{
		Users: users,
	}

	data, err := json.Marshal(userReponse)
	if err != nil {
		logging.Error.Print(common.FormatError(ctx, common.ServerError.ConvertJson, err))
		ServerErrorHandler(resp, req)
		return
	}

	resp.Write(data)
}

func (uc *UserController) RegisterUser(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	registerRequest := &registerUserRequest{}

	err := json.NewDecoder(req.Body).Decode(registerRequest)
	if err != nil {
		logging.Error.Printf(common.FormatError(ctx, common.ServerError.DecodeJson, err))
		ServerErrorHandler(resp, req)
		return
	}

	err = uc.UserService.RegisterUser(ctx, registerRequest.User)
	if err != nil {
		ServerErrorHandler(resp, req)
		return
	}

	CreatedHandler(resp, req, "User successfully registered!")
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}
