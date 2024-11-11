package router

import (
	"net/http"
	"github.com/gorilla/mux"
	"restapiv2/internal/http/itemsprocessor/handlers"
)

func NewItemsProcessorRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/stat", handlers.PrintStat)
	r.HandleFunc("/item/{key}", func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		key := vars["key"]
		handlers.GetItem(w, key)
	})
	r.HandleFunc("/item/{key}/{action}", func(w http.ResponseWriter, r *http.Request){})
	r.HandleFunc("/item/{key}/incr/{increment}", func(w http.ResponseWriter, r *http.Request){})
	r.Use(CountStat)
	r.Use(CheckMethods)
	return r
}