package handlers

import (
	"net/http"
	"restapiv2/internal/repository/itemscache"
	"fmt"
	"strconv"
)

func IncreaseItem(w http.ResponseWriter, key string, increment string) {

	val, exists := itemscache.Cache[key]

	if !exists {
		fmt.Fprint(w, "No such key in cache\n")
		return
	}
	
	currentVal := val
	
	currentValInt, err := strconv.Atoi(currentVal)
	if err != nil {
		fmt.Fprint(w, "Key value should be int\n")
		return
	}
	
	incInt, err := strconv.Atoi(increment)
	if err != nil {
		fmt.Fprint(w, "Increment value should be int\n")
		return
	}

	itemscache.Cache.UpdateItem(key, fmt.Sprintf("%d", currentValInt + incInt))
}


