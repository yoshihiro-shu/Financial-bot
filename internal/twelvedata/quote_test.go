package twelvedata_test

import (
	"testing"

	"github.com/yoshihiro-shu/financial-bot/internal/twelvedata"
)

var apiKey = ""

func TestQuote(t *testing.T) {
	if apiKey == "" {
		t.Skip()
	}
	client := twelvedata.NewClient(apiKey)
	quote, err := client.Quote("AAPL")
	if err != nil {
		t.Error(err)
	}
	if quote.Symbol == "" {
		t.Error("quote.Symbol == \"\"")
	}

	t.Logf("quote: %+v", quote)
}
