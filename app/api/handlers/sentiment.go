package handlers

import (
	"api/services"
	"encoding/json"
	"net/http"
)

type SentimentRequest struct {
	URL string `json:"url"`
}

type SentimentResponse struct {
	Sentiment string `json:"sentiment"`
}

func SentimentHandler(w http.ResponseWriter, r *http.Request) {
	var request SentimentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	if request.URL == "" {
		http.Error(w, "URL not provided", http.StatusBadRequest)
		return
	}

	sentiment, err := services.AnalyzeSentimentForURL(request.URL)
	if err != nil {
		http.Error(w, "Failed to analyze sentiment", http.StatusInternalServerError)
		return
	}

	responseData := SentimentResponse{Sentiment: sentiment}
	json.NewEncoder(w).Encode(responseData)
}
