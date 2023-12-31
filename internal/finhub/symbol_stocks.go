package finhub

import (
	"context"
	"time"

	"github.com/yoshihiro-shu/financial-bot/entity"
)

// Get real-time quote data for US stocks.
func (c *client) StockPriceBySymbol(symbol string) (*entity.Stock, error) {
	res, _, err := c.Quote(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		return nil, err
	}
	return &entity.Stock{
		Symbol: symbol,
		Open:   res.GetO(),
		Close:  res.GetC(),
		High:   res.GetH(),
		Low:    res.GetL(),
		Volume: 0,
		Date:   time.Now(),
	}, nil
}
