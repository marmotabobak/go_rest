package handlers

import (
	"net/http"
	"restapiv2/internal/repository/itemscache"
	"restapiv2/pkg/utils"
)

func PostItem(w http.ResponseWriter, cache *itemscache.Cache, key string, action string) {
	var currentVal, newVal string

	val, exists := cache.GetItem(key)
	if !exists {
		http.Error(w, "No such key in cache\n", http.StatusInternalServerError)
		return
	} else {
		currentVal = val
	}

	switch action {
	case "reverse":
		newVal = utils.ReversreString(currentVal)
	case "sort":
		newVal = utils.SortString(currentVal)
	case "dedup":
		newVal = utils.DeduplicateString(currentVal)
	default:
		http.Error(w, "No such action", http.StatusBadRequest)
	}

	cache.UpdateItem(key, newVal)

}
