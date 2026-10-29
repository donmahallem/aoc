package test_utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

type TestFileDb struct {
	basepath string
}

func (t *TestFileDb) CreateTestFilePath(year int, day int) string {
	return filepath.Join(t.basepath, fmt.Sprintf("../../test/data/full/%02d/%02d.txt", year, day))
}

func (t *TestFileDb) CheckTestDataExists(year int, day int) bool {
	return FileExists(t.CreateTestFilePath(year, day))
}

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
	return FileExists(createTestFilePath(year, day))
}

// PreloadFullData reads the full input for year and day into memory.
func PreloadFullData(year int, day int) ([]byte, error) {
	return os.ReadFile(createTestFilePath(year, day))
}

func GetTestData(year int, day int) (string, bool) {
	data, err := PreloadFullData(year, day)
	if err != nil {
		return "", false
	}
	return string(data), true
}

type Part int

const (
	Part1 Part = 1
	Part2 Part = 2
)

// TestFullDataForDate calls a function with the signature
// func(io.Reader) (A, error). If the function returns a non-nil error the
// test fails. This replaces the previous legacy helper that accepted
// func(io.Reader) A.
func TestFullDataForDate[A any](t *testing.T, year int, day int, fn func(in io.Reader) (A, error)) (A, bool) {
	t.Helper()
	input, err := PreloadFullData(year, day)
	return RunInput(t, requireInput(t, input, err, createTestFilePath(year, day)), fn), true
}

// requireInput skips the test when the file is missing and fails it on other read errors.
func requireInput(tb testing.TB, input []byte, err error, path string) []byte {
	tb.Helper()
	if errors.Is(err, fs.ErrNotExist) {
		tb.Skipf("Test data `%s` could not be found. Skipping...", path)
	}
	if err != nil {
		tb.Fatalf("Failed to read test data from path %s: %v", path, err)
	}
	return input
}

// RunInput calls fn on in-memory input and fails the test if fn returns an error.
func RunInput[A any](t *testing.T, input []byte, fn func(in io.Reader) (A, error)) A {
	t.Helper()
	res, err := fn(bytes.NewReader(input))
	if err != nil {
		t.Fatalf("function returned error: %v", err)
	}
	return res
}

// BenchInput times fn on in-memory input.
func BenchInput[A any](b *testing.B, input []byte, fn func(in io.Reader) (A, error)) {
	b.Helper()
	reader := bytes.NewReader(input)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		if _, err := fn(reader); err != nil {
			b.Fatalf("function returned error: %v", err)
		}
	}
}

func FileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	return err == nil && fileInfo.Mode().IsRegular()
}

func TestPartFromPath[A any](t *testing.T, path string, fn func(in io.Reader) (A, error)) (A, bool) {
	t.Helper()
	input, err := PreloadPartFromPath(path)
	return RunInput(t, requireInput(t, input, err, GetTestSampleFilePath(path)), fn), true
}

// PreloadPartFromPath reads a sample file, relative to test/, into memory.
func PreloadPartFromPath(path string) ([]byte, error) {
	return os.ReadFile(GetTestSampleFilePath(path))
}

/*
Method for testing full data in benchmarks
It will skip if the full test data file does not exist
*/
func BenchmarkFullDataForDate[A any](b *testing.B, year int, day int, fn func(in io.Reader) (A, error)) {
	b.Helper()
	input, err := PreloadFullData(year, day)
	BenchInput(b, requireInput(b, input, err, createTestFilePath(year, day)), fn)
}

func BenchmarkPartFromPath[A any](b *testing.B, path string, fn func(in io.Reader) (A, error)) {
	b.Helper()
	input, err := PreloadPartFromPath(path)
	BenchInput(b, requireInput(b, input, err, GetTestSampleFilePath(path)), fn)
}
