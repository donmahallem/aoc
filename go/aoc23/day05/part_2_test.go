package day05_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day05"
)

func TestPart2(t *testing.T) {
	result := day05.Part2(strings.NewReader(testData))
	if result != 46 {
		t.Errorf(`Expected %d to be %d`, result, 46)
	}
}

func BenchmarkPart2(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day05.Part2(reader)
	}
}
