package minkabu

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type Us struct{}

func (Us) PopularStocks() ([]string, error) {
	url := "https://us.minkabu.jp"

	var stocks []string
	c := colly.NewCollector()
	c.OnHTML(".w-laptop", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			code, err := extractStockCodeByUrl(e.Attr("href"))
			if err != nil {
				return
			}
			stocks = append(stocks, code)
		})
	})
	c.Visit(url)

	return stocks, nil
}

type Stock struct {
	Symbol string
	// 発行済株数
	OutstandingShares int64
	// 配当利回り
	DividendYield float32
	EPS           float32
	PER           float32
	PBR           float32
}

func (Us) StockDetail(code string) (*Stock, error) {
	url := fmt.Sprintf(UsStocksDetailURL, code)

	res := &Stock{
		Symbol: code,
	}
	c := colly.NewCollector()

	c.OnHTML(".flex-auto.bg-white", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			key := e.ChildText("th")
			value := e.ChildText("td")
			switch key {
			case "発行済株数":
				res.OutstandingShares = strToInt64(value)
			case "配当利回り":
				digits := findDecimalNumbers(value)
				res.DividendYield = strToFloat32(digits)
			case "EPS":
				digits := findDecimalNumbers(value)
				res.EPS = strToFloat32(digits)
			case "PER", "PER（調整後）":
				digits := findDecimalNumbers(value)
				res.PER = strToFloat32(digits)
			case "PBR":
				digits := findDecimalNumbers(value)
				res.PBR = strToFloat32(digits)
			}
		})
	})

	c.Visit(url)
	return res, nil
}
