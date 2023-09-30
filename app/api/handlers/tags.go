package handlers

import (
	"api/services"
	"encoding/json"
	"net/http"
)

func AvailableTagsHandler(w http.ResponseWriter, r *http.Request) {
	tags := services.GetDefaultTags()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}
