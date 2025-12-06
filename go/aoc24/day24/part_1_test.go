package day24_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day24"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

func TestPart1(t *testing.T) {
	t.Run("sample data", func(t *testing.T) {
		result := day24.Part1(strings.NewReader(testData))
		expected := uint64(2024)
		if result != expected {
			t.Errorf("Expected Part1 to be %v, got %v", expected, result)
		}
	})
	t.Run("actual data", func(t *testing.T) {

		result, ok := test_utils.TestFullDataForDate(t, 24, 24, day24.Part1)
		expected := uint64(65635066541798)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {

		reader := strings.NewReader(testData)
		for b.Loop() {
			day24.Part1(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 24, 24, day24.Part1)
	})
}
