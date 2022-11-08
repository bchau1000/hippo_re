package main

import (
	"fmt"
	"hippo/handler"
	"log"
	"net/http"
)

func main() {
	path := "/hippo/api/%s"

	http.HandleFunc(
		fmt.Sprintf(path, "version"),
		handler.GetVersion)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
