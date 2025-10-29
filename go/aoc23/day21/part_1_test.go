package day21_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day21"
)

var testDataSample1 string = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func TestParseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		reader := strings.NewReader(testDataSample1)
		res := day21.ParseInput(reader)
		if res.Width != 11 || res.Height != 11 {
			t.Errorf(`Expected width and height to be 11, got %d and %d`, res.Width, res.Height)
		}
		if res.StartX != 5 || res.StartY != 5 {
			t.Errorf(`Expected start position to be (5,5), got (%d,%d)`, res.StartX, res.StartY)
		}

	})

}

func TestPart1(t *testing.T) {
	t.Run("test sample 1", func(t *testing.T) {
		expected := 16
		reader := strings.NewReader(testDataSample1)
		res := day21.ParseInput(reader)
		visitedCount := day21.CountVisited(&res, 6, true)
		if visitedCount != expected {
			t.Errorf(`Expected number of blocks to be %d, got %d`, expected, visitedCount)
		}
	})
}

func BenchmarkPart1(b *testing.B) {

	reader := strings.NewReader(testDataSample1)
	for b.Loop() {
		day21.Part1(reader)
		reader.Seek(0, 0)
	}
}
