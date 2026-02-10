package day02_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day02"
)

func TestPart2(t *testing.T) {
	result, err := day02.Part2(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 4 {
		t.Errorf(`Expected %d to be %d`, result, 4)
	}
}

func BenchmarkPart2(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day02.Part2(reader)
	}
}
