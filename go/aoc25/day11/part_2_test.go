package day11_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc25/day11"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample2.txt
var testData2 string

func TestPart2(t *testing.T) {
	t.Run("test sample 1", func(t *testing.T) {
		expected := uint64(2)
		reader := strings.NewReader(testData2)
		result := day11.Part2(reader)
		if result != expected {
			t.Errorf(`Expected number of blocks to be %d, got %d`, expected, result)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		expected := uint64(367579641755680)
		result, ok := test_utils.TestFullDataForDate(t, 25, 11, day11.Part2)
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
		test_utils.BenchmarkFullDataForDate(b, 25, 11, day11.Part2)
	})
}
