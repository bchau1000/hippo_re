package errormsg

import (
	"context"
	"hippo/logging"

	"firebase.google.com/go/auth"
)

// Common messages for server fault errors
type ServerError string

// IDs for potential client fault errors
type ClientError string

const (
	// Server error details aren't sent to the client, only the ID
	DecodeJson    ServerError = "Error occurred while decoding JSON"
	ConvertJson   ServerError = "Error occurred while converting to JSON"
	ConvertSql    ServerError = "Error occurred while stringifying SQL"
	ExecuteSql    ServerError = "Error occurred while executing SQL"
	QueryFirebase ServerError = "Error occurred while querying Firebase"

	// Client errors are resolved on the client-side
	Custom       ClientError = "CLIENT_ERROR"
	EmailExists  ClientError = "FB_EMAIL_EXISTS"
	InvalidEmail ClientError = "FB_INVALID_EMAIL"
	UIDExists    ClientError = "FB_UID_EXISTS"
)

func FormatError(ctx context.Context, message ServerError, err error) string {
	return formatError(ctx, string(message), err)
}

func formatError(ctx context.Context, message string, err error) string {
	return logging.Errorf(ctx, string(message), err)
}

func IsClientError(err error) string {
	errorMessage := err.Error()

	if auth.IsEmailAlreadyExists(err) {
		return string(EmailExists)
	} else if auth.IsInvalidEmail(err) {
		return string(InvalidEmail)
	} else if auth.IsUIDAlreadyExists(err) {
		return string(UIDExists)
	} else if errorMessage == string(Custom) {
		return errorMessage
	}
	return ""
}
