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
	FirebaseRepository FirebaseRepository
}

func (ur *UserRepository) GetByIds(ctx context.Context) ([]model.User, error) {
	queryBuilder := sq.
		Select(
			ur.Column.Id,
			ur.Column.Email,
			ur.Column.DisplayName).
		From(ur.Table)

	var users []model.User

	rows, err := db.Search(ctx, queryBuilder)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := model.User{}

		err = rows.Scan(&user.Id, &user.Email, &user.DisplayName)
		if err != nil {
			logging.Error.Printf("Error encountered scanning user object: %v", err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) CreateUser(ctx context.Context, user model.UserToCreate) (int64, error) {
	var lastInsertId int64
	err := db.Transaction(ctx, func(tx *sql.Tx) error {
		newUser, err := ur.FirebaseRepository.RegisterUser(ctx, user)

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
		Table: "user",
		Column: userColumn{
			Uid:         "uid",
			Id:          "id",
			Email:       "email",
			DisplayName: "display_name",
		},
		FirebaseRepository: firebaseRepository,
	}
}
