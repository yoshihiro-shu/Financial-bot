package minkabu_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yoshihiro-shu/financial-bot/internal/scraping/minkabu"
)

func TestUsPopularStocks(t *testing.T) {
	us := minkabu.Us{}
	res, err := us.PopularStocks()
	if err != nil {
		t.Errorf("error is %s", err)
	}
	assert.Equal(t, 10, len(res))
	assert.NotEqual(t, "", res[0])
	assert.NotEqual(t, "", res[1])
	assert.NotEqual(t, "", res[2])
	assert.NotEqual(t, "", res[3])
	assert.NotEqual(t, "", res[4])
	assert.NotEqual(t, "", res[5])
	assert.NotEqual(t, "", res[6])
	assert.NotEqual(t, "", res[7])
	assert.NotEqual(t, "", res[8])
	assert.NotEqual(t, "", res[9])
}

func TestUsStockDetail(t *testing.T) {
	us := minkabu.Us{}
	res, err := us.StockDetail("AAPL")
	if err != nil {
		t.Errorf("error is %s", err)
	}
	assert.Equal(t, "AAPL", res.Symbol)
	assert.NotEqual(t, 0, res.OutstandingShares)
	assert.NotEqual(t, 0, res.DividendYield)
	assert.NotEqual(t, 0, res.EPS)
	assert.NotEqual(t, 0, res.PER)
	assert.NotEqual(t, 0, res.PBR)
}
