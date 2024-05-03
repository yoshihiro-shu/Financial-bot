package twelvedata

import (
	"encoding/json"
	"io"
	"net/http"
)

type responseEarnings struct {
	Meta struct {
		Symbol           string `json:"symbol"`
		Name             string `json:"name"`
		Currency         string `json:"currency"`
		Exchange         string `json:"exchange"`
		MicCode          string `json:"mic_code"`
		ExchangeTimezone string `json:"exchange_timezone"`
	} `json:"meta"`
	Earnings []struct {
		Date        string      `json:"date"`
		Time        string      `json:"time"`
		EpsEstimate interface{} `json:"eps_estimate"`
		EpsActual   interface{} `json:"eps_actual"`
		Difference  interface{} `json:"difference"`
		SurprisePrc interface{} `json:"surprise_prc"`
	} `json:"earnings"`
	Status string `json:"status"`
}

// https://api.twelvedata.com/earnings?symbol=AAPL&apikey=your_api_key
func (c *client) Earnings(symbol string) (*responseEarnings, error) {

	url := "https://api.twelvedata.com/earnings?symbol=" + symbol + "&apikey=" + c.apiKey
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var quote responseEarnings
	err = json.Unmarshal(responseData, &quote)
	if err != nil {
		return nil, err
	}
	return &quote, nil
}
