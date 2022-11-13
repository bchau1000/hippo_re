package service

import (
	"context"
	"hippo/repository"
)

type PingService struct {
	PingRepository repository.PingRepository
}

func (ps *PingService) PingDatabase(ctx context.Context) (bool, error) {
	success, err := ps.PingRepository.PingDatabase(ctx)
	return success, err
}

func NewPingService(pingRepository repository.PingRepository) PingService {
	return PingService{
		PingRepository: pingRepository,
	}
}
