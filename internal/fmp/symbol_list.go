package fmp

import (
	"io"
	"log"
	"net/http"

	"github.com/goccy/go-json"
)

type symbol struct {
	Symbol  string  `json:"symbol"`
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
	Ex      string  `json:"exchange"`
	ExShort string  `json:"exchangeShortName"`
	Type    string  `json:"type"`
}

// A list of all traded and non-traded stocks.
// https://site.financialmodelingprep.com/developer/docs/stock-market-quote-free-api/?direct=true
func (c *client) GetSymbolList() ([]symbol, error) {
	url := "https://financialmodelingprep.com/api/v3/stock/list?apikey=" + c.apiKey
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var symbolList []symbol
	err = json.Unmarshal(responseData, &symbolList)
	if err != nil {
		return nil, err
	}
	return symbolList, nil
}
