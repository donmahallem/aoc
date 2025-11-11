package day10_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day10"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestPart2(t *testing.T) {
	t.Run("test sample data", func(t *testing.T) {

		if result := day10.Part2(strings.NewReader(testData)); result != 81 {
			t.Errorf(`Expected %d to be %d`, result, 81)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		result, ok := test_utils.TestFullDataForDate(t, 24, 10, day10.Part2)
		if !ok || result != 1942 {
			t.Errorf(`Expected %d to be %d`, result, 1942)
		}
	})

}

func BenchmarkPart2(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day10.Part2(data)
	}
}
