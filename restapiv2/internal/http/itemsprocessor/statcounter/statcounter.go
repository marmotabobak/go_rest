package statcounter

import (
	"fmt"
	"net/http"
	"restapiv2/internal/repository/statstorage"

	"github.com/gorilla/mux"
)

type StatCounter struct {
	statStorage *statstorage.StatStorage
}

func NewStatCounter() *StatCounter {
	return &StatCounter{
		statStorage: statstorage.NewStatStorage(),
	}
}

func (sc *StatCounter) Count(handler http.Handler) http.Handler {
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

func (sc *StatCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, sc.statStorage.String())
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
