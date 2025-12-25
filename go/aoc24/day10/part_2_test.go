package day10_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day10"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestPart2(t *testing.T) {
	t.Run("test sample data", func(t *testing.T) {

		result, err := day10.Part2(strings.NewReader(testData))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != 81 {
			t.Errorf(`Expected %d to be %d`, result, 81)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		result, ok := test_utils.TestFullDataForDate(t, 24, 10, day10.Part2)
		if !ok || result != 1942 {
			t.Errorf(`Expected %d to be %d`, result, 1942)
		}
	})

}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		data := strings.NewReader(testData)
		for b.Loop() {
			data.Seek(0, io.SeekStart)
			day10.Part2(data)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 24, 10, day10.Part2)
	})
}
