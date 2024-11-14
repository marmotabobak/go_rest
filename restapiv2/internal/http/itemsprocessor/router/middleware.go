package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"restapiv2/internal/http/itemsprocessor/statcounter"
)

func CountStat(handler http.Handler) http.Handler {
	sc := statcounter.NewStatCounter()
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
