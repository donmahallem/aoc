package day15_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day15"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestPart2(t *testing.T) {

	t.Run("test sample data", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res, err := day15.Part2(reader)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if res != 145 {
			t.Errorf(`Expected number of blocks to be 145, got %d`, res)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		expected := uint64(269747)
		result, ok := test_utils.TestFullDataForDate(t, 23, 15, day15.Part2)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day15.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 15, day15.Part2)
	})
}
