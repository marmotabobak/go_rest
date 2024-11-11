package handlers

import (
	"net/http"
	"restapiv2/internal/repository/itemscache"
	"fmt"
)

func GetItem(w http.ResponseWriter, key string) {
	if val, exists := itemscache.Cache[key]; !exists {
		fmt.Fprint(w, "No such key in cache\n")
	} else {
		fmt.Fprintf(w, "%v\n", val) 
	}
}
