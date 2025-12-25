package day13_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day13"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

func TestPart1(t *testing.T) {

	t.Run("test sample data", func(t *testing.T) {
		reader := strings.NewReader(testData)
		res, err := day13.Part1(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res != 405 {
			t.Errorf(`Expected number of blocks to be 405, got %d`, res)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		expected := 42974
		result, ok := test_utils.TestFullDataForDate(t, 23, 13, day13.Part1)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day13.Part1(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 13, day13.Part1)
	})
}
