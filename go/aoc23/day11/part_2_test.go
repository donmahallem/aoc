package day11_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day11"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestPart2(t *testing.T) {
	t.Run("test sample data", func(t *testing.T) {
		const expected int = 82000210
		reader := strings.NewReader(testData)
		res, err := day11.Part2(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res != expected {
			t.Errorf(`Expected %d to be %d`, res, expected)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		expected := 630728425490
		result, ok := test_utils.TestFullDataForDate(t, 23, 11, day11.Part2)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day11.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 11, day11.Part2)
	})
}
