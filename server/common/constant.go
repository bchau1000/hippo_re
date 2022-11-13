package common

import (
	"context"
	"hippo/logging"
)

var Error = struct {
	DecodeJson  string
	ConvertJson string
	ConvertSql  string
	ExecuteSql  string
}{
	DecodeJson:  "Error occurred while decoding JSON",
	ConvertJson: "Error occurred while converting to JSON",
	ConvertSql:  "Error occurred while stringifying SQL",
	ExecuteSql:  "Error occurred while executing SQL",
}

func FormatError(ctx context.Context, message string, err error) string {
	return logging.Errorf(ctx, message, err)
}
