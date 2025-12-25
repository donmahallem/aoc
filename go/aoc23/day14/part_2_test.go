package day14_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day14"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestPart2(t *testing.T) {
	t.Run("test sample data", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res, err := day14.Part2(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res != 64 {
			t.Errorf(`Expected number of blocks to be 64, got %d`, res)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		expected := uint(79723)
		result, ok := test_utils.TestFullDataForDate(t, 23, 14, day14.Part2)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day14.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 14, day14.Part2)
	})
}
