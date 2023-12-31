package fmp_test

import (
	"os"
	"testing"

	"github.com/yoshihiro-shu/financial-bot/internal/fmp"
)

var apiKey = os.Getenv("FMP_API_KEY")

func TestGetSymbolList(t *testing.T) {
	client := fmp.NewClient(apiKey)
	symbolList, err := client.GetSymbolList()
	if err != nil {
		t.Error(err)
	}
	if len(symbolList) == 0 {
		t.Error("len(symbolList) == 0")
	}

	if symbolList[0].Symbol == "" {
		t.Error("symbolList[0].Symbol == \"\"")
	}

	if symbolList[0].Name == "" {
		t.Error("symbolList[0].Name == \"\"")
	}

	if symbolList[0].Price == 0 {
		t.Error("symbolList[0].Price == 0")
	}

	if symbolList[0].Ex == "" {
		t.Error("symbolList[0].Ex == \"\"")
	}

	if symbolList[0].ExShort == "" {
		t.Error("symbolList[0].ExShort == \"\"")
	}

	if symbolList[0].Type == "" {
		t.Error("symbolList[0].Type == \"\"")
	}
}
