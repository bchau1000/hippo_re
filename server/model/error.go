package model

import "fmt"

func Error(message string, errorGuid string) []byte {
	return []byte(fmt.Sprintf(
		"{\"message\": \"%s\", \"error\": \"%s\"}",
		message,
		errorGuid))
}
