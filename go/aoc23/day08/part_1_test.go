package day08_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day08"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample1.txt
var testData string

//go:embed sample2.txt
var testData2 string

func TestPart1(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 2
		reader := strings.NewReader(testData)
		res, err := day08.Part1(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res != expected {
			t.Errorf(`Expected %d to be %d`, res, expected)
		}
	})
	t.Run("testData2", func(t *testing.T) {
		const expected int = 6
		reader := strings.NewReader(testData2)
		res, err := day08.Part1(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res != expected {
			t.Errorf(`Expected %d to be %d`, res, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day08.Part1(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 8, day08.Part1)
	})
}
