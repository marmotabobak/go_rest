package handlers

import (
	"net/http"
	"restapiv2/internal/repository/itemscache"
	"strconv"
)

func IncreaseItem(w http.ResponseWriter, cache *itemscache.Cache, key string, increment string) {

	incInt, err := strconv.Atoi(increment)
	if err != nil {
		http.Error(w, "Increment value should be int\n", http.StatusBadRequest)
		return
	}

	cacheErr := cache.IncreaseValue(key, incInt)
	if cacheErr != nil {
		var httpStatus int
		switch cacheErr.Code {
		case itemscache.AbsentKeyErrorCode:
			httpStatus = http.StatusNotFound
		case itemscache.IncreaseErrorValueNotIntCode:
			httpStatus = http.StatusBadRequest
		default:
			httpStatus = http.StatusInternalServerError
		}
		http.Error(w, cacheErr.Desc, httpStatus)
		return		
	}
}
