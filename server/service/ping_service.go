package service

import (
	"hippo/repository"
)

type PingService struct {
	PingRepository repository.PingRepository
}

func (ps *PingService) PingDatabase() bool {
	success := ps.PingRepository.PingDatabase()
	return success
}

func NewPingService(pingRepository repository.PingRepository) PingService {
	return PingService{
		PingRepository: pingRepository,
	}
}
