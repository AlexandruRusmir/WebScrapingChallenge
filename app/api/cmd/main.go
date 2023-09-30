package main

import (
	"api/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/scrape", handlers.ScrapeHandler)
	http.HandleFunc("/analyzeSentiment", handlers.SentimentHandler)
	log.Fatal(http.ListenAndServe(":4200", nil))
}
