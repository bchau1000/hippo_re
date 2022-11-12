package service

import (
	"context"
	"hippo/repository"
)

type PingService struct {
	PingRepository repository.PingRepository
}

func (ps *PingService) PingDatabase(ctx context.Context) bool {
	success := ps.PingRepository.PingDatabase(ctx)
	return success
}

func NewPingService(pingRepository repository.PingRepository) PingService {
	return PingService{
		PingRepository: pingRepository,
	}
}
