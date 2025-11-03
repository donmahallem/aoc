package day06_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day06"
)

func TestPart2(t *testing.T) {
	result := day06.Part2(strings.NewReader(testData))
	if result != 71503 {
		t.Errorf(`Expected %d to be %d`, result, 71503)
	}
}

func BenchmarkPart2(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day06.Part2(reader)
	}
}
