package repository

import (
	"hippo/database"
	"hippo/logging"
)

type Repository struct {
	database           database.Database
	firebaseAuth       database.FirebaseAuth
	PingRepository     PingRepository
	UserRepository     UserRepository
	FirebaseRepository FirebaseRepository
}

func NewRepository(
	database database.Database,
	firebaseAuth database.FirebaseAuth,
) Repository {
	logging.Info.Print("Initializing repository dependencies")

	pingRepository := NewPingRepository(database, firebaseAuth)
	firebaseRepository := NewFirebaseRepository(database, firebaseAuth)
	userRepository := NewUserRepository(database, firebaseRepository)

	return Repository{
		database:           database,
		firebaseAuth:       firebaseAuth,
		PingRepository:     pingRepository,
		UserRepository:     userRepository,
		FirebaseRepository: firebaseRepository,
	}
}
