package day03_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day03"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

func TestPart1(t *testing.T) {
	t.Run("test sample data", func(t *testing.T) {
		result, err := day03.Part1(strings.NewReader(testData))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != 4484 {
			t.Errorf(`Expected %d to be %d`, result, 4484)
		}
	})
	t.Run("test sample data", func(t *testing.T) {
		result, err := day03.Part1(strings.NewReader("....@123\n456....."))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != 123 {
			t.Errorf(`Expected %d to be %d`, result, 123)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day03.Part1(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 3, day03.Part1)
	})
}
