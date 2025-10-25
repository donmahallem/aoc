package day18_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day18"
)

func TestPart2(t *testing.T) {
	t.Run("test block 2", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day18.Part2(reader)
		if res != 952408144115 {
			t.Errorf(`Expected number of blocks to be 952408144115, got %d`, res)
		}
	})
}

func BenchmarkPart2(b *testing.B) {

	reader := strings.NewReader(testData)
	for b.Loop() {
		day18.Part2(reader)
		reader.Seek(0, 0)
	}
}
