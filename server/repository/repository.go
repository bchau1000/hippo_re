package repository

import "hippo/logging"

type IRepository interface {
	TableName() string
}
type Repository struct {
	PingRepository PingRepository
	UserRepository UserRepository
}

func NewRepository() Repository {
	logging.Info.Print("Initializing repository dependencies")
	return Repository{
		PingRepository: NewPingRepository(),
		UserRepository: NewUserRepository(),
	}
}
