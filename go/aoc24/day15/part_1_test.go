package day15_test

import (
	_ "embed"
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day15"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample_1.txt
var testDataBig string

//go:embed sample_2.txt
var testDataSmall string

func TestPart1(t *testing.T) {
	t.Run("large", func(t *testing.T) {
		result, err := day15.Part1(strings.NewReader(testDataBig))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != 10092 {
			t.Errorf(`Expected %d to match 10092`, result)
		}
	})
	t.Run("small", func(t *testing.T) {
		result, err := day15.Part1(strings.NewReader(testDataSmall))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != 2028 {
			t.Errorf(`Expected %d to match 2028`, result)
		}
	})
	t.Run("test real data", func(t *testing.T) {
		result, ok := test_utils.TestFullDataForDate(t, 24, 15, day15.Part1)
		expected := 1430536
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		data := strings.NewReader(testDataBig)
		for b.Loop() {
			data.Seek(0, io.SeekStart)
			day15.Part1(data)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 24, 15, day15.Part1)
	})
}
