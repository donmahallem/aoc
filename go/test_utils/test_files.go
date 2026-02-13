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
	return filepath.Join(basepath, fmt.Sprintf("../../test/data/full/%02d/%02d.txt", year, day))
}

func getTestSampleFileBasePath() string {
	return "../../test/"
}

func GetTestSampleFilePath(path string) string {
	return filepath.Join(basepath, getTestSampleFileBasePath(), path)
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

// TestFullDataForDate calls a function with the signature
// func(io.Reader) (A, error). If the function returns a non-nil error the
// test fails. This replaces the previous legacy helper that accepted
// func(io.Reader) A.
func TestFullDataForDate[A any](t *testing.T, year int, day int, fn func(in io.Reader) (A, error)) (A, bool) {
	t.Helper()
	if CheckTestDataExists(year, day) {
		data, ok := GetTestData(year, day)
		if !ok {
			t.Fatalf("Failed to read test data for %d day %d", year, day)
		}
		res, err := fn(strings.NewReader(data))
		if err != nil {
			t.Fatalf("function returned error: %v", err)
		}
		return res, true
	} else {
		t.Skipf("Test data for %d day %d not found, skipping...", year, day)
	}
	var zero A
	return zero, false
}

func FileExists(path string) bool {
	if fileInfo, err := os.Stat(path); err == nil {
		return fileInfo.Mode().IsRegular()
	} else {
		return false
	}
}

func TestPartFromPath[A any](t *testing.T, path string, fn func(in io.Reader) (A, error)) (A, bool) {
	t.Helper()
	absoluePath := GetTestSampleFilePath(path)
	if FileExists(absoluePath) {
		data, err := os.ReadFile(absoluePath)
		if err != nil {
			t.Fatalf("Failed to read test data from path %s: %v", absoluePath, err)
		}
		res, err := fn(strings.NewReader(string(data)))
		if err != nil {
			t.Fatalf("function returned error: %v", err)
		}
		return res, true
	} else {
		t.Skipf("Test data file `%s` could not be found. Skipping...", absoluePath)
	}
	var zero A
	return zero, false
}

/*
Method for testing full data in benchmarks
It will skip if the full test data file does not exist
*/
func BenchmarkFullDataForDate[A any](b *testing.B, year int, day int, fn func(in io.Reader) (A, error)) {
	b.Helper() // define this function as a benchmark helper
	if CheckTestDataExists(year, day) {
		data, ok := GetTestData(year, day)
		if !ok {
			b.Fatalf("Failed to read test data for %d day %d", year, day)
		}
		reader := strings.NewReader(data)
		for b.Loop() {
			reader.Seek(0, io.SeekStart)
			_, err := fn(reader)
			if err != nil {
				b.Fatalf("function returned error: %v", err)
			}
		}
	} else {
		b.Skipf("Test data for %d day %d not found, skipping...", year, day)
	}
}

func BenchmarkPartFromPath[A any](b *testing.B, path string, fn func(in io.Reader) (A, error)) {
	b.Helper()
	absolutePath := GetTestSampleFilePath(path)
	if FileExists(absolutePath) {
		data, err := os.ReadFile(absolutePath)
		if err != nil {
			b.Fatalf("Failed to read test data from path %s: %v", absolutePath, err)
		}
		reader := strings.NewReader(string(data))
		for b.Loop() {
			reader.Seek(0, io.SeekStart)
			_, err := fn(reader)
			if err != nil {
				b.Fatalf("function returned error: %v", err)
			}
		}
	} else {
		b.Skipf("Test data file `%s` could not be found. Skipping...", absolutePath)
	}
}
