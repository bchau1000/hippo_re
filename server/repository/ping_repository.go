package repository

import (
	"context"
	db "hippo/database"

	sq "github.com/Masterminds/squirrel"
)

type PingRepository struct {
}

func (pr *PingRepository) PingDatabase(ctx context.Context) (bool, error) {
	query := sq.Select("1")
	_, err := db.Search(ctx, query)

	return err == nil, err
}

func NewPingRepository() PingRepository {
	return PingRepository{}
}
