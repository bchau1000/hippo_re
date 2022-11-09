package service

import (
	db "hippo/database"
	"hippo/logging"
)

type PingService struct {
}

func (ps *PingService) PingDatabase() {
	db.ExecuteQuery("SELECT 1;")
	logging.Info.Print("Successfully pinged the database")
}

func NewPingService() PingService {
	return PingService{}
}
