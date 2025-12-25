package day23_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day23"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestPart2(t *testing.T) {
	t.Run("test sample 1", func(t *testing.T) {
		expected := 154
		reader := strings.NewReader(testData)
		result, err := day23.Part2(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf(`Expected number of blocks to be %d, got %d`, expected, result)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		result, ok := test_utils.TestFullDataForDate(t, 23, 23, day23.Part2)
		if !ok || result != 6442 {
			t.Errorf(`Expected %d to be %d`, result, 6442)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day23.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 23, day23.Part2)
	})
}
