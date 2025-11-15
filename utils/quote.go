package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Quote struct {
	ID           string   `json:"_id"`
	Content      string   `json:"content"`
	Author       string   `json:"author"`
	Tags         []string `json:"tags"`
	AuthorSlug   string   `json:"authorSlug"`
	Length       int      `json:"length"`
	DateAdded    string   `json:"dateAdded"`
	DateModified string   `json:"dateModified"`
}

func GetRandomQuote() (*Quote, error) {
	resp, err := http.Get("http://api.quotable.io/quotes/random")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(body))
	var quotes []Quote

	err = json.Unmarshal(body, &quotes)
	if err != nil {
		return nil, err
	}
	return &quotes[0], nil
}
