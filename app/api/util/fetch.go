package util

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func FetchDocument(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("Failed to get content")
	}

	return goquery.NewDocumentFromReader(res.Body)
}
