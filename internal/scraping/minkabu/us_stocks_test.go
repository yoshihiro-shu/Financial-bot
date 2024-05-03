package minkabu_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yoshihiro-shu/financial-bot/internal/scraping/minkabu"
)

func TestUsPopularStocks(t *testing.T) {
	res, err := minkabu.UsPopularStocks()
	if err != nil {
		t.Errorf("error is %s", err)
	}
	assert.Equal(t, 10, len(res))
}
