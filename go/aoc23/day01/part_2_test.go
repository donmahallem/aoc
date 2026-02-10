package day01_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day01"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample_2.txt
var testDataPart2 string

func TestPart2(t *testing.T) {
	expected := 281
	reader := strings.NewReader(testDataPart2)
	res, err := day01.Part2(reader)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if res != expected {
		t.Errorf(`Expected %v to match %v`, res, expected)
	}
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day01.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 1, day01.Part2)
	})
}
