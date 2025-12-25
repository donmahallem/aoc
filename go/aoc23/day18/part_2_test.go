package day18_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day18"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestPart2(t *testing.T) {
	t.Run("test sample data", func(t *testing.T) {
		res, err := day18.Part2(strings.NewReader(testData))
		if err != nil {
			t.Errorf("Part2() error = %v", err)
			return
		}
		if res != 952408144115 {
			t.Errorf("Part2() = %v, want %v", res, 952408144115)
		}
	})
	t.Run("test real data", func(t *testing.T) {
		expected := int64(131431655002266)
		result, ok := test_utils.TestFullDataForDate(t, 23, 18, day18.Part2)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day18.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 18, day18.Part2)
	})
}
