package day06_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day06"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestPart2(t *testing.T) {
	result, err := day06.Part2(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 71503 {
		t.Errorf(`Expected %d to be %d`, result, 71503)
	}
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day06.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 6, day06.Part2)
	})
}
