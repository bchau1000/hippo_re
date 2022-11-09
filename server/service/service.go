package service

type Service struct {
	PingService PingService
}

func NewService() Service {
	return Service{
		PingService: NewPingService(),
	}
}
