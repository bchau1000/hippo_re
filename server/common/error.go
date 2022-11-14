package common

import (
	"context"
	"hippo/logging"

	"firebase.google.com/go/auth"
)

// Common messages for server fault errors
var ServerError = struct {
	DecodeJson    string
	ConvertJson   string
	ConvertSql    string
	ExecuteSql    string
	QueryFirebase string
}{
	DecodeJson:    "Error occurred while decoding JSON",
	ConvertJson:   "Error occurred while converting to JSON",
	ConvertSql:    "Error occurred while stringifying SQL",
	ExecuteSql:    "Error occurred while executing SQL",
	QueryFirebase: "Error occurred while querying Firebase",
}

var ClientError = struct {
}{}

func FormatError(ctx context.Context, message string, err error) string {
	return logging.Errorf(ctx, message, err)
}

func IsClientError(err error) bool {
	return auth.IsEmailAlreadyExists(err) ||
		auth.IsInvalidEmail(err) ||
		auth.IsUIDAlreadyExists(err)
}
