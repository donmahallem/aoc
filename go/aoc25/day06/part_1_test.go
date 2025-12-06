package day06_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc25/day06"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

func TestPart1(t *testing.T) {
	t.Run("test sample 1", func(t *testing.T) {
		expected := uint64(4277556)
		reader := strings.NewReader(testData)
		result := day06.Part1(reader)
		if result != expected {
			t.Errorf(`Expected number of blocks to be %d, got %d`, expected, result)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		expected := uint64(4583860641327)
		result, ok := test_utils.TestFullDataForDate(t, 25, 6, day06.Part1)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {

		reader := strings.NewReader(testData)
		for b.Loop() {
			day06.Part1(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 25, 6, day06.Part1)
	})
}
