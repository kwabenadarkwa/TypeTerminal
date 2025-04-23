package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
)

type Quotes struct {
	Quotes []Quote
}

func (q *Quotes) len() int {
	return len(q.Quotes)
}

type Quote struct {
	Id     int    `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func GetAllQuotes(path string) Quotes {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var quotes Quotes

	json.Unmarshal(byteValue, &quotes.Quotes)

	//INFO: testings
	// for i := range quotes.Quotes {
	// 	fmt.Println("this happens")
	// 	fmt.Println("Quote Id: " + strconv.Itoa(quotes.Quotes[i].Id))
	// 	fmt.Println("Actual Quote: " + quotes.Quotes[i].Quote)
	// 	fmt.Println("Author: " + quotes.Quotes[i].Author)
	// }
	//
	defer jsonFile.Close()

	return quotes
}

func SelectRandomQuoteFromQuotes(quotes Quotes) Quote {
	max := quotes.len()
	randomIndex := rand.IntN(max)
	return quotes.Quotes[randomIndex]
}
