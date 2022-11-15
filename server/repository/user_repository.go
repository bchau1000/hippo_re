package repository

import (
	"context"
	"database/sql"
	"hippo/common/errormsg"
	db "hippo/database"
	"hippo/logging"
	"hippo/model"

	sq "github.com/Masterminds/squirrel"
)

type userColumn struct {
	Id          string `db:"id"`
	Uid         string `db:"uid"`
	Email       string `db:"email"`
	DisplayName string `db:"display_name"`
}

// Repository for the `user` table
type UserRepository struct {
	Table              string     // table name
	Column             userColumn // columns
	firebaseRepository FirebaseRepository
}

func (ur *UserRepository) AuthUser(ctx context.Context, idToken string) (model.User, error) {
	return ur.firebaseRepository.AuthUser(ctx, idToken)
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	return ur.firebaseRepository.GetUserByEmail(ctx, email)
}

func (ur *UserRepository) CreateUser(ctx context.Context, user model.UserToCreate) (int64, error) {
	var lastInsertId int64
	err := db.Transaction(ctx, func(tx *sql.Tx) error {
		newUser, err := ur.firebaseRepository.RegisterUser(ctx, user)

		if err != nil {
			return err
		}

		userToCreate := sq.
			Insert(ur.Table).
			Columns(
				ur.Column.Uid,
				ur.Column.Email,
				ur.Column.DisplayName).
			Values(
				newUser.UID,
				newUser.Email,
				newUser.DisplayName)

		result, err := db.InsertTx(ctx, tx, userToCreate)
		if err != nil {
			return err
		}

		lastInsertId, err = result.LastInsertId()
		if err != nil {
			logging.Error.Print(errormsg.FormatError(ctx, "Error occurred while parsing last inserted ID", err))
			return err
		}

		return nil
	})

	if err != nil {
		return -1, err
	}

	return lastInsertId, nil
}

func NewUserRepository(firebaseRepository FirebaseRepository) UserRepository {
	return UserRepository{
		// Private
		firebaseRepository: firebaseRepository,

		// Public
		Table: "user",
		Column: userColumn{
			Uid:         "uid",
			Id:          "id",
			Email:       "email",
			DisplayName: "display_name",
		},
	}
}
