package handlers

import (
	"net/http"
	"restapiv2/internal/repository/itemscache"
)

func Deleteitem(w http.ResponseWriter, cache *itemscache.Cache, key string) {
	_, exists := cache.GetItem(key)
	if !exists {
		http.Error(w, "No such key in cache\n", http.StatusInternalServerError)
		return
	}
	cache.DeleteItem(key)
}
