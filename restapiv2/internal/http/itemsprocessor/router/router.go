package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"restapiv2/internal/http/itemsprocessor/handlers"
)

type ItemsProcessorRouter struct {
	MuxRouter *mux.Router
}

func NewItemsProcessorRouter() *ItemsProcessorRouter {
	r := mux.NewRouter()
	r.HandleFunc("/stat", StatHandler)
	r.HandleFunc("/item/{key}", GetItemHandler)
	r.HandleFunc("/item/{key}/{action}", PostHandler)
	r.HandleFunc("/item/{key}/incr/{increment}", Increasehandler)
	r.Use(CountStat)
	
	return &ItemsProcessorRouter {
		MuxRouter: r,
	}
}

func (ipr *ItemsProcessorRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ipr.MuxRouter.ServeHTTP(w, r)
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.PrintStat(w)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	switch r.Method {
	case http.MethodGet:
		handlers.GetItem(w, key)
	case http.MethodPut:
		handlers.PutItem(w, r, key)
	case http.MethodDelete:
		handlers.Deleteitem(w, key)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		vars := mux.Vars(r)
		key := vars["key"]
		action := vars["action"]

		if action == "reverse" || action == "sort" || action == "dedup" {
			handlers.PostItem(w, action, key)
		} else {
			http.Error(w, "Unknown action", http.StatusBadRequest)
		}
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func Increasehandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		vars := mux.Vars(r)
		key := vars["key"]
		increment := vars["increment"]
		handlers.IncreaseItem(w, key, increment)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
