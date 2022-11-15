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

func (fr *FirebaseRepository) RegisterUser(ctx context.Context, request model.UserToCreate) (model.User, error) {
	userToCreateFb := &auth.UserToCreate{}
	userToCreateFb.
		Email(request.User.Email).
		Password(request.Password).
		DisplayName(request.User.DisplayName)

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

func (fr *FirebaseRepository) AuthUser(ctx context.Context, idToken string) (model.User, error) {
	authToken, err := database.GetAuth().VerifyIDToken(ctx, idToken)
	if err != nil {
		logging.Error.Print(
			errormsg.FormatError(
				ctx,
				errormsg.QueryFirebase,
				err))
	}

	user, err := fr.GetUserByUID(ctx, authToken.UID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (fr *FirebaseRepository) GetUserByUID(ctx context.Context, uid string) (model.User, error) {
	userRecord, err := database.GetAuth().GetUser(ctx, uid)
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

func (fr *FirebaseRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
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
