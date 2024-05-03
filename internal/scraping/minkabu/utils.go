package minkabu

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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

func findDecimalNumbers(input string) string {
	re := regexp.MustCompile(`\d+\.\d+`)
	return re.FindAllString(input, -1)[0]
}

func strToInt64(str string) int64 {
	cleanedNumber := strings.ReplaceAll(str, ",", "")
	i, err := strconv.ParseInt(cleanedNumber, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func strToFloat32(str string) float32 {
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0
	}
	return float32(f)
}
