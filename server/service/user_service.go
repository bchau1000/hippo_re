package service

import (
	"context"
	"hippo/model"
	"hippo/repository"
)

type UserService struct {
	UserRepository     repository.UserRepository
	FirebaseRepository repository.FirebaseRepository
}

func (us *UserService) GetByIds(ctx context.Context) ([]model.User, error) {
	users, err := us.UserRepository.GetByIds(ctx)
	return users, err
}

func (us *UserService) GetByEmail(ctx context.Context, email string) (model.User, error) {
	user, err := us.FirebaseRepository.GetUsers(ctx, email)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (us *UserService) RegisterUser(ctx context.Context, userToCreate model.UserToCreate) error {
	_, err := us.UserRepository.CreateUser(ctx, userToCreate)
	if err != nil {
		return err
	}

	return err
}

func NewUserService(
	userRepository repository.UserRepository,
	firebaseRepository repository.FirebaseRepository) UserService {

	return UserService{
		UserRepository:     userRepository,
		FirebaseRepository: firebaseRepository,
	}
}
