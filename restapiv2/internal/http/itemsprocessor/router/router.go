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

	r.Handle("/stat", statCountHandler).Methods(http.MethodGet)
	r.HandleFunc("/item/{key}", i.getItemHandler).Methods(http.MethodGet, http.MethodDelete, http.MethodPut)
	r.HandleFunc("/item/{key}/{action}", i.postItemHandler).Methods(http.MethodPost)
	r.HandleFunc("/item/{key}/incr/{increment}", i.increaseItemHandler).Methods(http.MethodPost)
	r.Use(statCountHandler.Count)

	return &i
}

func (i *ItemsProcessorRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i.muxRouter.ServeHTTP(w, r)
}

// TODO: Нужно ли выносить логику маршрутизации в хэндлеры-функции или оставить здесь в маршрутизаторе?
// Для StatCounter вынес в него, но он является не функцией, а полноценным хэндлером 

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
	}
}

func (i *ItemsProcessorRouter) postItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	action := vars["action"]

	if action == "reverse" || action == "sort" || action == "dedup" {
		handlers.PostItem(w, i.cache, key, action)
	} else {
		http.Error(w, "Unknown action", http.StatusBadRequest)
	}
}

func (i *ItemsProcessorRouter) increaseItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	increment := vars["increment"]
	handlers.IncreaseItem(w, i.cache, key, increment)
}
