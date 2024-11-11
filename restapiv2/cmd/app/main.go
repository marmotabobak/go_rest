package main

import (
	"log"
	"net/http"
	"restapiv2/internal/http/itemsprocessor/router"
	"restapiv2/internal/repository/stat"
)

func main() {
	stat.Init()
	r := router.NewItemsProcessorRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}