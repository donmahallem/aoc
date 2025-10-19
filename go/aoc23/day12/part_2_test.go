package day12_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day12"
)

func TestPart2(t *testing.T) {
	t.Run("testData", func(t *testing.T) {
		reader := strings.NewReader(testData)
		result := day12.Part2(reader)
		const expectedLen int = 525152
		if result != expectedLen {
			t.Errorf(`Expected %d arrangements to be found, got %d`, expectedLen, result)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day12.Part2(reader)
	}
}
