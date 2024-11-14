package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"restapiv2/internal/repository/statstorage"
)

type StatCounter struct {
	statStorage *statstorage.StatStorageType
}

func NewStatCounter() *StatCounter {
	return &StatCounter{
		statStorage: &statstorage.StatStorage,
	}
}

func (sc *StatCounter) CountStat(handler http.Handler) http.Handler {
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
