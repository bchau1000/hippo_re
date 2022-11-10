package controller

import (
	"encoding/json"
	"hippo/model"
	"hippo/service"
	"net/http"
)

type UserResponse struct {
	Users []model.User `json:"users"`
}

type UserController struct {
	UserService service.UserService
}

func (uc *UserController) GetUsers(resp http.ResponseWriter, request *http.Request) {
	users, err := uc.UserService.GetByIds()
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write(model.Error("Error occurred while fulfilling request", "REQUEST"))
		return
	}

	userReponse := UserResponse{
		Users: users,
	}

	data, err := json.Marshal(userReponse)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write(model.Error("Error occurred while decoding json", "JSON"))
		return
	}

	resp.Write(data)
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}
