package config

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func InitFirebase(cfg *Config) (*firebase.App, error) {
	opt := option.WithCredentialsFile(cfg.Firebase.Path)
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return app, nil
}
