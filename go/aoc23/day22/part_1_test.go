package day22_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day22"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample_1.txt
var testDataSample1 string

func TestPart1(t *testing.T) {
	t.Run("test sample 1", func(t *testing.T) {
		expected := 5
		reader := strings.NewReader(testDataSample1)
		result, err := day22.Part1(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf(`Expected number of blocks to be %d, got %d`, expected, result)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		result, ok := test_utils.TestFullDataForDate(t, 23, 22, day22.Part1)
		if !ok || result != 503 {
			t.Errorf(`Expected %d to be %d`, result, 503)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testDataSample1)
		for b.Loop() {
			day22.Part1(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 22, day22.Part1)
	})
}
