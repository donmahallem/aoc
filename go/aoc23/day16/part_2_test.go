package day16_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day16"
)

func TestPart2(t *testing.T) {
	t.Run("test block 2", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day16.Part2(reader)
		if res != 46 {
			t.Errorf(`Expected number of blocks to be 46, got %d`, res)
		}
	})
}

func BenchmarkPart2(b *testing.B) {

	reader := strings.NewReader(testData)
	for b.Loop() {
		day16.Part2(reader)
		reader.Seek(0, 0)
	}
}
