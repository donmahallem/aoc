package day16_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day16"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestPart2(t *testing.T) {
	t.Run("test sample data", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res, err := day16.Part2(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res != 51 {
			t.Errorf(`Expected number of blocks to be 51, got %d`, res)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		expected := 7330
		result, ok := test_utils.TestFullDataForDate(t, 23, 16, day16.Part2)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day16.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 16, day16.Part2)
	})
}
