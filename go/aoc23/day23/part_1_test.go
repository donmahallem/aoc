package day23_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day23"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

func TestPart1(t *testing.T) {
	t.Run("test sample 1", func(t *testing.T) {
		expected := 94
		reader := strings.NewReader(testData)
		result := day23.Part1(reader)
		if result != expected {
			t.Errorf(`Expected number of blocks to be %d, got %d`, expected, result)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		result, ok := test_utils.TestFullDataForDate(t, 23, 23, day23.Part1)
		if !ok || result != 2094 {
			t.Errorf(`Expected %d to be %d`, result, 2094)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample 1", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			reader.Seek(0, 0)
			day23.Part1(reader)
		}
	})

	b.Run("benchmark real data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 23, day23.Part1)
	})
}
