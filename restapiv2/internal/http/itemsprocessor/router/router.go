package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"restapiv2/internal/http/itemsprocessor/handlers"
	"restapiv2/internal/repository/itemscache"
)

type ItemsProcessorRouter struct {
	muxRouter *mux.Router
	cache *itemscache.Cache
}

func NewItemsProcessorRouter() *ItemsProcessorRouter {

	cache := itemscache.NewCache()
	statCountHandler := handlers.NewStatCountHandler()
	r := mux.NewRouter()

	i := ItemsProcessorRouter{
		muxRouter: r,
		cache: cache,
	}

	r.Handle("/stat", statCountHandler)
	r.HandleFunc("/item/{key}", i.getItemHandler)
	// r.HandleFunc("/item/{key}/{action}", postHandler)
	// r.HandleFunc("/item/{key}/incr/{increment}", increasehandler)
	r.Use(statCountHandler.Count)

	return &i

}

func (i *ItemsProcessorRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i.muxRouter.ServeHTTP(w, r)
}

// TODO: Нужно ли выносить логику маршрутизации в хэндлеры-функции или оставить здесь в маршрутизаторе?
// Для StatCounter вынес в него, но и он является не функцией, а хэндлером полноценным

func (i *ItemsProcessorRouter) getItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	switch r.Method {
	case http.MethodGet:
		handlers.GetItem(w, i.cache, key)
	case http.MethodPut:
		handlers.PutItem(w, r, i.cache, key)
	case http.MethodDelete:
		handlers.Deleteitem(w, i.cache, key)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// func postHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		vars := mux.Vars(r)
// 		key := vars["key"]
// 		action := vars["action"]

// 		if action == "reverse" || action == "sort" || action == "dedup" {
// 			handlers.PostItem(w, action, key)
// 		} else {
// 			http.Error(w, "Unknown action", http.StatusBadRequest)
// 		}
// 	} else {
// 		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
// 	}
// }

// func increasehandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		vars := mux.Vars(r)
// 		key := vars["key"]
// 		increment := vars["increment"]
// 		handlers.IncreaseItem(w, key, increment)
// 	} else {
// 		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
// 	}
// }
