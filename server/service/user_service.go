package service

import (
	"hippo/model"
	"hippo/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (us *UserService) GetByIds() ([]model.User, error) {
	users, err := us.UserRepository.GetByIds()
	return users, err
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{
		UserRepository: userRepository,
	}
}
