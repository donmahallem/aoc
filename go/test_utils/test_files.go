package test_utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
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

func TestFullDataForDate[A any](t *testing.T, year int, day int, fn func(in io.Reader) A) (A, bool) {
	t.Helper() // define this function as a test helper
	if CheckTestDataExists(year, day) {
		data, ok := GetTestData(year, day)
		if !ok {
			t.Fatalf("Failed to read test data for %d day %d", year, day)
		}
		return fn(strings.NewReader(data)), true
	} else {
		t.Skipf("Test data for %d day %d not found, skipping...", year, day)
	}
	var zero A
	return zero, false
}

/*
Method for testing full data in benchmarks
It will skip if the full test data file does not exist
*/
func BenchmarkFullDataForDate[A any](b *testing.B, year int, day int, fn func(in io.Reader) A) {
	b.Helper() // define this function as a benchmark helper
	if CheckTestDataExists(year, day) {
		data, ok := GetTestData(year, day)
		if !ok {
			b.Fatalf("Failed to read test data for %d day %d", year, day)
		}
		reader := strings.NewReader(data)
		for b.Loop() {
			reader.Seek(0, io.SeekStart)
			fn(reader)
		}
	} else {
		b.Skipf("Test data for %d day %d not found, skipping...", year, day)
	}
}
