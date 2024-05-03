package twelvedata

import (
	"encoding/json"
	"io"
	"net/http"
)

type responseQuote struct {
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	Exchange      string `json:"exchange"`
	MicCode       string `json:"mic_code"`
	Currency      string `json:"currency"`
	Datetime      string `json:"datetime"`
	Timestamp     int    `json:"timestamp"`
	Open          string `json:"open"`
	High          string `json:"high"`
	Low           string `json:"low"`
	Close         string `json:"close"`
	Volume        string `json:"volume"`
	PreviousClose string `json:"previous_close"`
	Change        string `json:"change"`
	PercentChange string `json:"percent_change"`
	AverageVolume string `json:"average_volume"`
	IsMarketOpen  bool   `json:"is_market_open"`
	FiftyTwoWeek  struct {
		Low               string `json:"low"`
		High              string `json:"high"`
		LowChange         string `json:"low_change"`
		HighChange        string `json:"high_change"`
		LowChangePercent  string `json:"low_change_percent"`
		HighChangePercent string `json:"high_change_percent"`
		Range             string `json:"range"`
	} `json:"fifty_two_week"`
}

// https://twelvedata.com/docs#quote
func (c *client) Quote(symbol string) (*responseQuote, error) {

	url := "https://api.twelvedata.com/quote?symbol=" + symbol + "&apikey=" + c.apiKey
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var quote responseQuote
	err = json.Unmarshal(responseData, &quote)
	if err != nil {
		return nil, err
	}
	return &quote, nil
}
