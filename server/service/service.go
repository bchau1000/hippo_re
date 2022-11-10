package service

import (
	"hippo/logging"
	"hippo/repository"
)

type Service struct {
	PingService PingService
}

func NewService(repository repository.Repository) Service {
	logging.Info.Print("Initializing service dependencies")
	return Service{
		PingService: NewPingService(repository.PingRepository),
	}
}
