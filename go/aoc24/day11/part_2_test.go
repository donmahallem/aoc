package day11_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day11"
)

func TestPart2(t *testing.T) {
	if result := day11.Part2(strings.NewReader(testData)); result != 65601038650482 {
		t.Errorf(`Expected %d to contain %d`, result, 65601038650482)
	}
}

func BenchmarkPart2(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day11.Part2(data)
	}
}
