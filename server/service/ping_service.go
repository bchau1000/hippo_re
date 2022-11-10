package service

import (
	"hippo/logging"
	"hippo/repository"
)

type PingService struct {
	PingRepository repository.PingRepository
}

func (ps *PingService) PingDatabase() bool {
	success := ps.PingRepository.PingDatabase()

	if success {
		logging.Info.Print("Successfully pinged the database")
	}

	return success
}

func NewPingService(pingRepository repository.PingRepository) PingService {
	return PingService{
		PingRepository: pingRepository,
	}
}
