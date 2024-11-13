package handlers

import (
	"fmt"
	"net/http"
	"restapiv2/internal/repository/itemscache"
)

func GetItem(w http.ResponseWriter, key string) {
	val, exists := itemscache.Cache[key]
	if !exists {
		http.Error(w, "No such key in cache\n", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%v\n", val)
}
