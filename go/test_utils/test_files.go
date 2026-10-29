package test_utils

import (
	"bytes"
	"errors"
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

// samplePath resolves path relative to test/.
func samplePath(path string) string {
	return filepath.Join(basepath, "../../test", path)
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

// runInput calls fn on in-memory input and fails the test if fn returns an error.
func runInput[A any](t *testing.T, input []byte, fn func(in io.Reader) (A, error)) A {
	t.Helper()
	res, err := fn(bytes.NewReader(input))
	if err != nil {
		t.Fatalf("function returned error: %v", err)
	}
	return res
}

// benchInput times fn on in-memory input.
func benchInput[A any](b *testing.B, input []byte, fn func(in io.Reader) (A, error)) {
	b.Helper()
	reader := bytes.NewReader(input)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		if _, err := fn(reader); err != nil {
			b.Fatalf("function returned error: %v", err)
		}
	}
}

func TestPartFromPath[A any](t *testing.T, path string, fn func(in io.Reader) (A, error)) (A, bool) {
	t.Helper()
	file := samplePath(path)
	input, err := os.ReadFile(file)
	return runInput(t, requireInput(t, input, err, file), fn), true
}

func BenchmarkPartFromPath[A any](b *testing.B, path string, fn func(in io.Reader) (A, error)) {
	b.Helper()
	file := samplePath(path)
	input, err := os.ReadFile(file)
	benchInput(b, requireInput(b, input, err, file), fn)
}
