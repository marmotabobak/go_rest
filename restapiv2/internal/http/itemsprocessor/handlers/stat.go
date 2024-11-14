package handlers

import (
	"fmt"
	"net/http"
	"restapiv2/internal/repository/statstorage"
)

func PrintStat(w http.ResponseWriter) {
	fmt.Fprint(w, statstorage.StatStorage.String())
}
