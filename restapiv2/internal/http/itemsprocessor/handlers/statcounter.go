package handlers

import (
	"fmt"
	"net/http"
	"restapiv2/internal/repository/statstorage"
	"github.com/gorilla/mux"
)

type StatCountHandler struct {
	statStorage *statstorage.StatStorage
}

func NewStatCountHandler() *StatCountHandler {
	return &StatCountHandler{
		statStorage: statstorage.NewStatStorage(),
	}
}

func (sc *StatCountHandler) Count(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var processedAction string
		vars := mux.Vars(r)
		_, itemIsDefined := vars["item"]
		action, actionIsDefined := vars["action"]
		_, incrIsDefined := vars["incr"]

		if r.URL.Path == "/stat" {
			processedAction = "stat"
		}

		if itemIsDefined && !(actionIsDefined || incrIsDefined) {
			processedAction = "item"
		}

		if actionIsDefined {
			processedAction = action + " item"
		}

		if incrIsDefined {
			processedAction = "increase item"
		}

		statAction := fmt.Sprintf("%s %s", r.Method, processedAction)
		sc.statStorage.Update(statAction)

		handler.ServeHTTP(w, r)
	})
}

func (sc *StatCountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, sc.statStorage.String())
}
