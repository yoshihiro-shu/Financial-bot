package testutils

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
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
