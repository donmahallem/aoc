package day21_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day21"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample_1.txt
var testDataSample1 string

func TestPart1(t *testing.T) {

	t.Run("test sample data", func(t *testing.T) {
		reader := strings.NewReader(testDataSample1)
		res, err := day21.Part1(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := 42
		if res != expected {
			t.Errorf(`Expected number of blocks to be %d, got %d`, expected, res)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		expected := 3697
		result, ok := test_utils.TestFullDataForDate(t, 23, 21, day21.Part1)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testDataSample1)
		for b.Loop() {
			day21.Part1(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 21, day21.Part1)
	})
}
