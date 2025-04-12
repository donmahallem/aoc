package test_utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func createTestFilePath(year int, day int) string {
	return filepath.Join(basepath, fmt.Sprintf("../../data/%02d/%02d.txt", year, day))
}

func CheckTestDataExists(year int, day int) bool {
	finalPath := createTestFilePath(year, day)
	if fileInfo, err := os.Stat(finalPath); err == nil {
		return fileInfo.Mode().IsRegular()
	} else {
		return false
	}
}

func GetTestData(year int, day int) (string, bool) {
	file, err := os.ReadFile(createTestFilePath(year, day))
	if err != nil {
		return "", false
	} else {
		return string(file), true
	}
}
