package day05_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day05"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

func TestPart1(t *testing.T) {
	result, err := day05.Part1(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 35 {
		t.Errorf(`Expected %d to be %d`, result, 35)
	}
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day05.Part1(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 5, day05.Part1)
	})
}
