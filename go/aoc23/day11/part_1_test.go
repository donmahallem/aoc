package day11_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day11"
)

var testData string = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestPart1(t *testing.T) {
	t.Run("parseInput", func(t *testing.T) {
		t.Run("offset 1", func(t *testing.T) {
			reader := strings.NewReader(testData)
			galaxies := day11.ParseInput(reader, 1)
			const expectedLen int = 9
			if len(galaxies) != expectedLen {
				t.Errorf(`Expected %d galaxies to be parsed, got %d`, expectedLen, len(galaxies))
			}
		})
		t.Run("offset 100", func(t *testing.T) {
			reader := strings.NewReader(testData)
			galaxies := day11.ParseInput(reader, 100)
			const expectedLen int = 9
			if len(galaxies) != expectedLen {
				t.Errorf(`Expected %d galaxies to be parsed, got %d`, expectedLen, len(galaxies))
			}
		})
	})
	t.Run("testData1", func(t *testing.T) {
		const expected int = 374
		reader := strings.NewReader(testData)
		if res := day11.Part1(reader); res != expected {
			t.Errorf(`Expected %d to be %d`, res, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day11.Part1(reader)
	}
}
