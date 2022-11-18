package database

import (
	"context"
	"hippo/config"
	"hippo/logging"

	"firebase.google.com/go/auth"
)

type FirebaseAuth struct {
	Client *auth.Client
}

func NewFirebaseAuth(ctx context.Context, cfg *config.Config) FirebaseAuth {
	authClient, err := cfg.Firebase.App.Auth(ctx)
	if err != nil {
		logging.Error.Fatalf("Fatal error occurred while getting Firebase Authentication: %v", err)
	}
	logging.Info.Print("Successfully opened Firebase Authentication connection")

	return FirebaseAuth{
		Client: authClient,
	}
}
