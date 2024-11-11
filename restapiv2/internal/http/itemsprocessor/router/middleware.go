package router

import (
	"net/http"
	"restapiv2/pkg/utils"
	"github.com/gorilla/mux"
)

func CheckMethods(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_, actionIsDefined := vars["action"]
		_, incrIsDefined := vars["incr"]

		if r.Method == http.MethodPost && !(actionIsDefined || incrIsDefined) {
			http.Error(w, "action should be defined for POST method", http.StatusMethodNotAllowed)
			return
		}

		if utils.MethodIsGetPutDelete(r.Method) && (actionIsDefined || incrIsDefined) {
			http.Error(w, "GET/PUT/DELETE methods are not allowed for actions", http.StatusMethodNotAllowed)
			return
		}

		if r.URL.Path == "/stat" && utils.MethodIsPutDelete(r.Method) {
			http.Error(w, "stat action has only GET method", http.StatusMethodNotAllowed)
			return			
		}

		handler.ServeHTTP(w, r)
	})
}
