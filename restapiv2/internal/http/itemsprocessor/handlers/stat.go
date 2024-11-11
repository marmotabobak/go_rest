package handlers

import (
	"fmt"
	"net/http"
	"restapiv2/internal/repository/stat"
)

func PrintStat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, stat.StatStorage.String())
}
