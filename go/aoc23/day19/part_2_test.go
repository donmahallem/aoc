package day19_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day19"
)

func TestPart2(t *testing.T) {
	t.Run("test block 2", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day19.Part2(reader)
		if res != 167409079868000 {
			t.Errorf(`Expected number of blocks to be 167409079868000, got %d`, res)
		}
	})
}

func BenchmarkPart2(b *testing.B) {

	reader := strings.NewReader(testData)
	for b.Loop() {
		day19.Part2(reader)
		reader.Seek(0, 0)
	}
}
