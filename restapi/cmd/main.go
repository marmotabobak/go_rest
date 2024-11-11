package main

import (
	"fmt"
	"log"
	"mityamentor/cmd/application/restapi"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "Hello from urvantsev")
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	}

	http.HandleFunc("/hello-yandex", handler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func startAPI() {
	restapi.StartAPI()
}
