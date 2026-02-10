package day07_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day07"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestPart2(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		var expected int = 5905
		reader := strings.NewReader(testData)
		res, err := day07.Part2(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res != expected {
			t.Errorf(`Expected %v to match %v`, res, expected)
		}
	})
	t.Run("testData3", func(t *testing.T) {
		var expected int = 9800
		reader := strings.NewReader(testData3)
		res, err := day07.Part2(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res != expected {
			t.Errorf(`Expected %v to match %v`, res, expected)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day07.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 7, day07.Part2)
	})
}
