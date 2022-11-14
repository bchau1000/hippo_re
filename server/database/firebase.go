package database

import (
	"context"
	"hippo/config"
	"hippo/logging"

	"firebase.google.com/go/auth"
)

var authClient *auth.Client

func GetAuth() *auth.Client {
	return authClient
}

func InitFirebase(ctx context.Context, cfg *config.Config) {
	logging.Info.Print("Initializing Firebase Authentication")

	var err error
	authClient, err = cfg.Firebase.App.Auth(ctx)
	if err != nil {
		logging.Error.Fatalf("Fatal error occurred while getting Firebase Authentication: %v", err)
	}
}
