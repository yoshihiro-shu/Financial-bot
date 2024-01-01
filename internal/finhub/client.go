package finhub

import finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"

type client struct {
	*finnhub.DefaultApiService
}

func NewClient(apiKey string) *client {
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", apiKey)
	return &client{finnhub.NewAPIClient(cfg).DefaultApi}
}
