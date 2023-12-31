package finhub

import (
	"context"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

type news struct {
	Category string `json:"category"` // News category.
	Datetime int64  `json:"datetime"` // Published time in UNIX timestamp.
	Headline string `json:"headline"` // News headline.
	Id       int64  `json:"id"`       // News ID. This value can be used for minId params to get the latest news only.
	Image    string `json:"image"`    // Thumbnail image URL.
	Related  string `json:"related"`  // Related stocks and companies mentioned in the article.
	Source   string `json:"source"`   // News source.
	Summary  string `json:"summary"`  // News summary.
	Url      string `json:"url"`      // URL of the original article.
}

// Get latest market news. https://finnhub.io/docs/api/market-news
func (c *client) GetMarketNews() ([]finnhub.MarketNews, error) {
	res, _, err := c.MarketNews(context.Background()).Category("general").Execute()
	if err != nil {
		return nil, err
	}
	return res, nil
}
