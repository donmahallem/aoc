package day03_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day03"
)

func TestPart2(t *testing.T) {
	result := day03.Part2(strings.NewReader(testData))
	if result != 467835 {
		t.Errorf(`Expected %d to be %d`, result, 467835)
	}
}

func BenchmarkPart2(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day03.Part2(reader)
	}
}
