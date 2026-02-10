package day07_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day07"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

const testData2 string = `32T3K 100
43T4Q 200`

//go:embed sample2.txt
var testData3 string

func TestPart1(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		var expected int = 6440
		reader := strings.NewReader(testData)
		res, err := day07.Part1(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res != expected {
			t.Errorf(`Expected %v to match %v`, res, expected)
		}
	})
	t.Run("testData2", func(t *testing.T) {
		var expected int = 500
		reader := strings.NewReader(testData2)
		res, err := day07.Part1(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res != expected {
			t.Errorf(`Expected %v to match %v`, res, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day07.Part1(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 7, day07.Part1)
	})
}
