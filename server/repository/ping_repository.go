package repository

import (
	db "hippo/database"

	sq "github.com/Masterminds/squirrel"
)

type PingRepository struct {
}

func (pr *PingRepository) PingDatabase() bool {
	query := sq.Select("1")
	_, err := db.Search(query)

	return err == nil
}

func NewPingRepository() PingRepository {
	return PingRepository{}
}
