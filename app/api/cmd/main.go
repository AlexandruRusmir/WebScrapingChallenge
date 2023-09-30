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
	log.Fatal(http.ListenAndServe(":4200", nil))
}
