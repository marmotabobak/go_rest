package main

import (
	"log"
	"net/http"
	"restapiv2/internal/http/itemsprocessor/router"
)

func main() {
	router := router.NewItemsProcessorRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
