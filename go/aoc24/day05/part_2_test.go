package day05_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day05"
)

func TestPart2(t *testing.T) {
	result, err := day05.Part2(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 123 {
		t.Errorf(`Expected %d to be %d`, result, 123)
	}
}

func BenchmarkPart2(b *testing.B) {
	testData := strings.NewReader(testData)
	for b.Loop() {
		testData.Seek(0, io.SeekStart)
		day05.Part2(testData)
	}
}
