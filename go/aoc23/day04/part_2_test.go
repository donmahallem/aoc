package day04_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day04"
)

func TestPart2(t *testing.T) {
	result := day04.Part2(strings.NewReader(testData))
	if result != 30 {
		t.Errorf(`Expected winners to have a length of %d. Not %d`, 30, result)
	}
}

func BenchmarkPart2(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day04.Part2(reader)
	}
}
