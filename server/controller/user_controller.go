package controller

import (
	"encoding/json"
	constant "hippo/common"
	"hippo/logging"
	"hippo/model"
	"hippo/service"
	"net/http"
)

type userResponse struct {
	Users []model.User `json:"users"`
}

type UserController struct {
	UserService service.UserService
}

func (uc *UserController) GetUsers(resp http.ResponseWriter, req *http.Request) {
	users, err := uc.UserService.GetByIds(req.Context())
	if err != nil {
		// Error is propagated -- no need to log
		ServerErrorHandler(resp, req)
		return
	}

	userReponse := userResponse{
		Users: users,
	}

	data, err := json.Marshal(userReponse)
	if err != nil {
		logging.Error.Printf(
			logging.Errorf(
				req.Context(),
				constant.Error.DecodeJson,
				err))
		ServerErrorHandler(resp, req)
		return
	}

	resp.Write(data)
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}
