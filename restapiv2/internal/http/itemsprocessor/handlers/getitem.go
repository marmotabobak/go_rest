package handlers

import (
	"net/http"
	"restapiv2/internal/repository/itemscache"
	"fmt"
)

func GetItem(w http.ResponseWriter, key string) {

	val, exists := itemscache.Cache[key]

	if !exists {
		fmt.Fprint(w, "No such key in cache\n")
		return
	}

	fmt.Fprintf(w, "%v\n", val) 
}