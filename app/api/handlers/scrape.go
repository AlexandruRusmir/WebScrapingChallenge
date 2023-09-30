package handlers

import (
	"api/services"
	"encoding/json"
	"net/http"
)

type ScrapeRequest struct {
	URL  string   `json:"url"`
	Tags []string `json:"tags,omitempty"`
}

func ScrapeHandler(w http.ResponseWriter, r *http.Request) {
	var req ScrapeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := services.ScrapeContentWithWordCount(req.URL, req.Tags)
	if err != nil {
		http.Error(w, "Failed to scrape content", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}
