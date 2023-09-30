package main

import (
	"api/handlers"
	"api/middlewares"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/scrape", middlewares.EnableCORS(handlers.ScrapeHandler))
	http.HandleFunc("/analyzeSentiment", middlewares.EnableCORS(handlers.SentimentHandler))
	http.HandleFunc("/availableTags", middlewares.EnableCORS(handlers.AvailableTagsHandler))
	log.Fatal(http.ListenAndServe(":4200", nil))
}
