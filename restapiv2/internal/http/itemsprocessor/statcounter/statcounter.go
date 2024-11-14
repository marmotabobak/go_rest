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

func (sc *StatCounter) Update(itemAction string) {
	sc.statStorage.Update(itemAction)
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
		sc.Update(statAction)

		handler.ServeHTTP(w, r)
	})
}

func (sc *StatCounter) PrintStat(w http.ResponseWriter) {
	fmt.Fprint(w, sc.statStorage.String())
	// fmt.Fprint(w, "STAT")
}
