package testutils

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// GetProjectRoot returns the root directory of the project.
func GetProjectRoot() (string, error) {
	// 現在のファイルの絶対パスを取得
	_, b, _, _ := runtime.Caller(0)
	path := filepath.Dir(b)

	index := strings.LastIndex(path, "financial-bot")
	if index == -1 {
		fmt.Println("Path does not contain 'financial-bot'")
		return "", fmt.Errorf("path does not contain 'financial-bot'")
	}

	// "financial-bot"を含むパスを取得
	result := path[:index+len("financial-bot")]
	return result, nil
}

// ParseUnixTime parses unix time
func ParseUnixTime(unixTime string) time.Time {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalf("Failed to load location: %s", err)
	}
	const layout = "2006-01-02 15:04:05 -0700"
	res, err := time.Parse(layout, unixTime)
	if err != nil {
		log.Fatalf("Failed to parse time: %s", err)
	}

	return res.In(location)
}
