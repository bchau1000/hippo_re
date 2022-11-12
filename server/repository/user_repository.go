package repository

import (
	"context"
	db "hippo/database"
	"hippo/logging"
	"hippo/model"

	sq "github.com/Masterminds/squirrel"
)

type userColumn struct {
	Id          string
	Email       string
	DisplayName string
}

// Repository for the `user` table
type UserRepository struct {
	Table  string     // table name
	Column userColumn // columns
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

func NewUserRepository() UserRepository {
	return UserRepository{
		Table: "user",
		Column: userColumn{
			Id:          "id",
			Email:       "email",
			DisplayName: "display_name",
		},
	}
}
