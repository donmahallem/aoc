package day10_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day10"
	"github.com/donmahallem/aoc/go/test_utils"
)

const testData string = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestPart1(t *testing.T) {
	t.Run("test sample data", func(t *testing.T) {

		if result := day10.Part1(strings.NewReader(testData)); result != 36 {
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
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day10.Part1(data)
	}
}
