package handlers

import (
	"fmt"
	"net/http"
	"restapiv2/internal/repository/stat"
)

func PrintStat(w http.ResponseWriter) {
	fmt.Fprint(w, stat.StatStorage.String())
}
