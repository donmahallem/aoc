package day17_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day17"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestPart2(t *testing.T) {
	t.Run("sample", func(t *testing.T) {
		res, err := day17.Part2(strings.NewReader(testData))
		if err != nil {
			t.Errorf("Part2() error = %v", err)
			return
		}
		if res != 94 {
			t.Errorf("Part2() = %v, want %v", res, 94)
		}
	})
	t.Run("sample2", func(t *testing.T) {
		res, err := day17.Part2(strings.NewReader(`111111111111
999999999991
999999999991
999999999991
999999999991`))
		if err != nil {
			t.Errorf("Part2() error = %v", err)
			return
		}
		if res != 71 {
			t.Errorf("Part2() = %v, want %v", res, 71)
		}
	})
	t.Run("test real data", func(t *testing.T) {
		expected := uint32(1101)
		result, ok := test_utils.TestFullDataForDate(t, 23, 17, day17.Part2)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day17.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 17, day17.Part2)
	})
}
