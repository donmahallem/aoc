package day08_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day08"
	"github.com/donmahallem/aoc/go/test_utils"
)

var testData3 string = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestPart2(t *testing.T) {
	t.Run("testData3", func(t *testing.T) {
		const expected uint = 6
		reader := strings.NewReader(testData3)
		res, err := day08.Part2(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res != expected {
			t.Errorf(`Expected %d to be %d`, res, expected)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day08.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 8, day08.Part2)
	})
}
