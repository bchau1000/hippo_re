package repository

import (
	"context"
	"hippo/common/errormsg"
	"hippo/database"
	"hippo/logging"
	"hippo/model"

	"firebase.google.com/go/auth"
)

type FirebaseRepository struct {
}

func (fr *FirebaseRepository) RegisterUser(ctx context.Context, user model.UserToCreate) (model.User, error) {
	userToCreateFb := &auth.UserToCreate{}
	userToCreateFb.
		Email(user.Email).
		Password(user.Password).
		DisplayName(user.DisplayName)

	userRecord, err := database.
		GetAuth().
		CreateUser(ctx, userToCreateFb)
	if err != nil {
		logging.Error.Print(
			errormsg.FormatError(
				ctx,
				errormsg.QueryFirebase,
				err))
		return model.User{}, err
	}

	return *model.UserRecordToUser(userRecord), nil

}

func (fr *FirebaseRepository) GetUsers(ctx context.Context, email string) (model.User, error) {
	user, err := database.GetAuth().GetUserByEmail(ctx, email)
	if err != nil {
		logging.Error.Print(
			errormsg.FormatError(
				ctx,
				errormsg.QueryFirebase,
				err))
		return model.User{}, err
	}

	return model.User{
		Id:          1,
		Email:       user.Email,
		DisplayName: user.DisplayName,
	}, nil
}

func NewFirebaseRepository() FirebaseRepository {
	return FirebaseRepository{}
}
