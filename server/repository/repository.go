package repository

import (
	"hippo/logging"
)

type IRepository interface {
	TableName() string
}
type Repository struct {
	PingRepository     PingRepository
	UserRepository     UserRepository
	FirebaseRepository FirebaseRepository
}

func NewRepository() Repository {
	logging.Info.Print("Initializing repository dependencies")

	firebaseRepository := NewFirebaseRepository()
	pingRepository := NewPingRepository()
	userRepository := NewUserRepository(firebaseRepository)

	return Repository{
		PingRepository:     pingRepository,
		UserRepository:     userRepository,
		FirebaseRepository: firebaseRepository,
	}
}
