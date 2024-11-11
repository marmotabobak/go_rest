package main

import (
	"log"
	"net/http"
	"restapiv2/internal/http/itemsprocessor/router"
	"restapiv2/internal/repository/stat"
	"restapiv2/internal/repository/itemscache"
)

func main() {
	stat.Init()
	itemscache.Init()
	r := router.NewItemsProcessorRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}