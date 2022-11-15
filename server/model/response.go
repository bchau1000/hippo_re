package model

import "fmt"

type Response struct{}

func Error(message string, errorGuid string) []byte {
	return []byte(fmt.Sprintf(
		"{\"message\": \"%s\", \"error\": \"%s\"}",
		message,
		errorGuid))
}

func Success(message string) []byte {
	return []byte(fmt.Sprintf(
		"{\"message\": \"%s\"}",
		message))
}
