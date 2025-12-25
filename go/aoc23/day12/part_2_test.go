package day12_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day12"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestPart2(t *testing.T) {
	t.Run("test sample data", func(t *testing.T) {
		reader := strings.NewReader(testData)
		result, err := day12.Part2(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		const expectedLen int = 525152
		if result != expectedLen {
			t.Errorf(`Expected %d arrangements to be found, got %d`, expectedLen, result)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		expected := 4964259839627
		result, ok := test_utils.TestFullDataForDate(t, 23, 12, day12.Part2)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day12.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 12, day12.Part2)
	})
}
