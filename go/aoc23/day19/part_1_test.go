package day19_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day19"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

func TestPart1(t *testing.T) {
	t.Run("test sample data", func(t *testing.T) {
		res, err := day19.Part1(strings.NewReader(testData))
		if err != nil {
			t.Errorf("Part1() error = %v", err)
			return
		}
		if res != 19114 {
			t.Errorf("Part1() = %v, want %v", res, 19114)
		}
	})
	t.Run("test real data", func(t *testing.T) {
		expected := 319295
		result, ok := test_utils.TestFullDataForDate(t, 23, 19, day19.Part1)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day19.Part1(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 19, day19.Part1)
	})
}
