package day01_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day01"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample_1.txt
var testData string

func TestPart1(t *testing.T) {
	expected := 142
	reader := strings.NewReader(testData)
	res, err := day01.Part1(reader)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if res != expected {
		t.Errorf(`Expected %v to match %v`, res, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day01.Part1(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 1, day01.Part1)
	})
}
