package handlers

import (
	"fmt"
	"net/http"
	"restapiv2/internal/repository/itemscache"
)

func GetItem (w http.ResponseWriter, cache *itemscache.Cache, key string) {
	val, exists := cache.GetItem(key)
	if !exists {
		http.Error(w, "No such key in cache\n", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%v\n", val)
}
