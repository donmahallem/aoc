package day15_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day15"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

const testDataHash string = `HASH`

func TestPart1(t *testing.T) {
	t.Run("test sample 1", func(t *testing.T) {

		reader := strings.NewReader(testDataHash)
		res, err := day15.Part1(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res != 52 {
			t.Errorf(`Expected number of blocks to be 52, got %d`, res)
		}
	})
	t.Run("test sample 2", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res, err := day15.Part1(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res != 1320 {
			t.Errorf(`Expected number of blocks to be 1320, got %d`, res)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		expected := uint(507769)
		result, ok := test_utils.TestFullDataForDate(t, 23, 15, day15.Part1)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		reader := strings.NewReader(testData)
		for b.Loop() {
			day15.Part1(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 23, 15, day15.Part1)
	})
}
