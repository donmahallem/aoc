package day10_test

import (
	_ "embed"
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day10"
	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

func TestPart1(t *testing.T) {
	t.Run("test sample data", func(t *testing.T) {

		result, err := day10.Part1(strings.NewReader(testData))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != 36 {
			t.Errorf(`Expected %d to contain %d`, result, 36)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		result, ok := test_utils.TestFullDataForDate(t, 24, 10, day10.Part1)
		if !ok || result != 796 {
			t.Errorf(`Expected %d to be %d`, result, 796)
		}
	})

}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		data := strings.NewReader(testData)
		for b.Loop() {
			data.Seek(0, io.SeekStart)
			day10.Part1(data)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 24, 10, day10.Part1)
	})
}
