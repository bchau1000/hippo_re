package repository

import "hippo/logging"

type Repository struct {
	PingRepository PingRepository
}

func NewRepository() Repository {
	logging.Info.Print("Initializing repository dependencies")
	return Repository{
		PingRepository: NewPingRepository(),
	}
}
