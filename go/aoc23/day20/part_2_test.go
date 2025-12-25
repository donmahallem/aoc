package day20_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day20"
	"github.com/donmahallem/aoc/go/test_utils"
)

const testOutputId int = ('o' << 40) + ('u' << 32) + ('t' << 24) + ('p' << 16) + ('u' << 8) + 't'

func TestPart2(t *testing.T) {
	t.Run("test sample 2", func(t *testing.T) {
		expected := 1
		reader := strings.NewReader(testDataSample2)
		res := day20.HandlePart2(reader, testOutputId)
		if res != expected {
			t.Errorf(`Expected number of blocks to be %d, got %d`, expected, res)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testDataSample1)
		for b.Loop() {
			day20.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 20, day20.Part2)
	})
}
