package service

import (
	"context"
	"hippo/model"
	"hippo/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (us *UserService) GetByIds(ctx context.Context) ([]model.User, error) {
	users, err := us.UserRepository.GetByIds(ctx)
	return users, err
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{
		UserRepository: userRepository,
	}
}
