package minkabu_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yoshihiro-shu/financial-bot/internal/scraping/minkabu"
)

func TestJPStockDetail(t *testing.T) {
	jp := minkabu.Jp{}
	res, err := jp.StockDetail("9434")
	if err != nil {
		t.Errorf("error is %s", err)
	}
	assert.Equal(t, "9434", res.Symbol)
	// assert.NotEqual(t, 0, res.OutstandingShares)
	assert.NotEqual(t, 0, res.DividendYield)
	assert.NotEqual(t, 0, res.PSR)
	assert.NotEqual(t, 0, res.PER)
	assert.NotEqual(t, 0, res.PBR)
}
