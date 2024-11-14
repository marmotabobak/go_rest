package handlers

// import (
// 	"fmt"
// 	"net/http"
// 	"restapiv2/internal/repository/itemscache"
// 	"strconv"
// )

// func IncreaseItem(w http.ResponseWriter, key string, increment string) {

// 	val, exists := itemscache.Cache[key]

// 	if !exists {
// 		http.Error(w, "No such key in cache\n", http.StatusInternalServerError)
// 		return
// 	}

// 	currentVal := val

// 	currentValInt, err := strconv.Atoi(currentVal)
// 	if err != nil {
// 		http.Error(w, "Key value should be int\n", http.StatusInternalServerError)
// 		return
// 	}

// 	incInt, err := strconv.Atoi(increment)
// 	if err != nil {
// 		http.Error(w, "Increment value should be int\n", http.StatusInternalServerError)
// 		return
// 	}

// 	itemscache.Cache.UpdateItem(key, fmt.Sprintf("%d", currentValInt+incInt))
// }
