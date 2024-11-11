package router

import (
	"net/http"
	"github.com/gorilla/mux"
)

func NewItemsProcessorRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/stat", func(w http.ResponseWriter, r *http.Request){})
	r.HandleFunc("/item/{key}", func(w http.ResponseWriter, r *http.Request){})
	r.HandleFunc("/item/{key}/{action}", func(w http.ResponseWriter, r *http.Request){})
	r.HandleFunc("/item/{key}/incr/{increment}", func(w http.ResponseWriter, r *http.Request){})
	r.Use(CheckMethods)
	return r
}