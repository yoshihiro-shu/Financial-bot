package minkabu

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly/v2"
)

func UsPopularStocks() ([]string, error) {
	url := "https://us.minkabu.jp"

	var stocks []string
	c := colly.NewCollector()
	c.OnHTML(".w-laptop", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			code, err := extractStockCodeByUrl(e.Attr("href"))
			fmt.Println(code)
			if err != nil {
				return
			}
			stocks = append(stocks, code)
		})
	})
	c.Visit(url)

	return stocks, nil
}

// Extract stock code from relative url
// /stocks/NVDA -> NVDA
func extractStockCodeByUrl(url string) (string, error) {
	regex, err := regexp.Compile(`/stocks/([A-Z]+)`)
	if err != nil {
		return "", err
	}
	code := regex.FindAllStringSubmatch(url, -1)[0][1]
	if code == "" {
		return "", fmt.Errorf("code is empty")
	}
	return code, nil
}
