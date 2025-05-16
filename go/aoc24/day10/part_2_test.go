package day10_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day10"
)

func TestPart2(t *testing.T) {
	if result := day10.Part2(strings.NewReader(testData)); result != 81 {
		t.Errorf(`Expected %d to contain %d`, result, 81)
	}
}

func BenchmarkPart2(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day10.Part2(data)
	}
}
