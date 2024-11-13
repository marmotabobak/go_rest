package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"restapiv2/internal/models"
	"restapiv2/internal/repository/itemscache"
)

func PutItem(w http.ResponseWriter, r *http.Request, key string) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error while parsing request body\n", http.StatusInternalServerError)
		return
	}
	defer func() { r.Body.Close() }()

	var item models.Item
	err = json.Unmarshal(body, &item)
	if err != nil {
		http.Error(w, "Error while parsing request body to json\n", http.StatusInternalServerError)
		return
	}

	itemscache.Cache.UpdateItem(key, item.Data.Value)
}
