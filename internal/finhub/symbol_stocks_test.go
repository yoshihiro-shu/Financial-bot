package finhub_test

import (
	"os"
	"testing"

	"github.com/yoshihiro-shu/financial-bot/internal/finhub"
)

var apiKey = os.Getenv("FINHUB_API_KEY")

func TestSymbolStocks(t *testing.T) {
	client := finhub.NewClient(apiKey)
	stock, err := client.SymbolStocks("AAPL")
	if err != nil {
		t.Errorf("error is %s", err)
	}

	if stock.High == 0 {
		t.Errorf("stock is %v", stock)
	}
	if stock.Low == 0 {
		t.Errorf("stock is %v", stock)
	}
	if stock.Open == 0 {
		t.Errorf("stock is %v", stock)
	}
	if stock.Close == 0 {
		t.Errorf("stock is %v", stock)
	}
	// Volumeは取得できない
	// if stock.Volume == 0 {
	// 	t.Errorf("stock is %v", stock)
	// }
	if stock.Date.IsZero() {
		t.Errorf("stock is %v", stock)
	}
}
