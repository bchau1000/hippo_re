package service

import (
	"context"
	"hippo/model"
	"hippo/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (us *UserService) GetByEmail(ctx context.Context, email string) (model.User, error) {
	return us.UserRepository.GetUserByEmail(ctx, email)
}

func (us *UserService) AuthUser(ctx context.Context, idToken string) (model.User, error) {
	return us.UserRepository.AuthUser(ctx, idToken)
}

func (us *UserService) RegisterUser(ctx context.Context, userToCreate model.UserToCreate) error {
	_, err := us.UserRepository.CreateUser(ctx, userToCreate)
	return err
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{
		UserRepository: userRepository,
	}
}
