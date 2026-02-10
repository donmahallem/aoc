package day19_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day19"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestPart2(t *testing.T) {
	t.Run("test sample data", func(t *testing.T) {
		res, err := day19.Part2(strings.NewReader(testData))
		if err != nil {
			t.Errorf("Part2() error = %v", err)
			return
		}
		expected := 167409079868000
		if res != expected {
			t.Errorf("Part2() = %d, want %d", res, expected)
		}
	})
	t.Run("test real data", func(t *testing.T) {
		expected := 110807725108076
		result, ok := test_utils.TestFullDataForDate(t, 23, 19, day19.Part2)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day19.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 19, day19.Part2)
	})
}
