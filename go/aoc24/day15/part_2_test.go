package day15_test

import (
	_ "embed"
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day15"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestPart2(t *testing.T) {
	t.Run("test sample data small", func(t *testing.T) {

		result, err := day15.Part2(strings.NewReader(testDataSmall))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := 1751
		if result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
	t.Run("test sample data", func(t *testing.T) {

		result, err := day15.Part2(strings.NewReader(testDataBig))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := 9021
		if result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		result, ok := test_utils.TestFullDataForDate(t, 24, 15, day15.Part2)
		expected := 1452348
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})

}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		data := strings.NewReader(testDataBig)
		for b.Loop() {
			data.Seek(0, io.SeekStart)
			day15.Part2(data)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 24, 15, day15.Part2)
	})
}
