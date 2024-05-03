package minkabu

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type JPStock struct {
	Symbol string
	// OutstandingShares int64
	DividendYield float32
	PER           float32
	PSR           float32
	PBR           float32
}

type Jp struct{}

func (Jp) StockDetail(code string) (*JPStock, error) {
	url := fmt.Sprintf(JpStockDetailURL, code)

	res := &JPStock{
		Symbol: code,
	}
	c := colly.NewCollector()

	c.OnHTML(".md_box", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			key := e.ChildText("th")
			value := e.ChildText("td")
			switch key {
			// 4,756,200千株という書き方に対応する
			// case "発行済株数":
			// 	res.OutstandingShares = strToInt64(value)
			case "配当利回り":
				digits := findDecimalNumbers(value)
				res.DividendYield = strToFloat32(digits)
			case "PSR":
				digits := findDecimalNumbers(value)
				res.PSR = strToFloat32(digits)
			case "PER(調整後)":
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
