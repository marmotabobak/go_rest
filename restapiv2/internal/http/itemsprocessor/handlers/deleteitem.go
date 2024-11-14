package handlers

// import (
// 	"net/http"
// 	"restapiv2/internal/repository/itemscache"
// )

// func Deleteitem(w http.ResponseWriter, key string) {
// 	if _, exists := itemscache.Cache[key]; !exists {
// 		http.Error(w, "No such key in cache\n", http.StatusInternalServerError)
// 		return
// 	}
// 	itemscache.Cache.DeleteItem(key)
// }
