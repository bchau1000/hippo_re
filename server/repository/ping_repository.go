package repository

import (
	"context"
	"hippo/database"
)

type PingRepository struct {
	database database.Database
	fbAuth   database.FirebaseAuth
}

func (pr *PingRepository) PingDatabase(ctx context.Context) (bool, error) {
	err := pr.database.Ping(ctx)
	if err != nil {
		return false, err
	}

	return err == nil, err
}

func NewPingRepository(
	database database.Database,
	fbAuth database.FirebaseAuth,
) PingRepository {
	return PingRepository{
		database: database,
		fbAuth:   fbAuth,
	}
}
