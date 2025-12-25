package day17_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day17"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

func TestPart1(t *testing.T) {
	t.Run("test sample data", func(t *testing.T) {
		res, err := day17.Part1(strings.NewReader(testData))
		if err != nil {
			t.Errorf("Part1() error = %v", err)
			return
		}
		if res != 102 {
			t.Errorf("Part1() = %v, want %v", res, 102)
		}
	})
	t.Run("test real data", func(t *testing.T) {
		expected := uint32(967)
		result, ok := test_utils.TestFullDataForDate(t, 23, 17, day17.Part1)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day17.Part1(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 17, day17.Part1)
	})
}
