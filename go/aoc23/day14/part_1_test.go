package day14_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day14"
)

var testData string = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func TestParseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 10
		reader := strings.NewReader(testData)
		if res := day14.ParseInput(reader); len(res) != expected {
			t.Errorf(`Expected %d to be %d`, len(res), expected)
		}

	})
	t.Run("test block 1", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day14.ParseInput(reader)
		expectedRows := []uint{
			5,
			2,
			4,
			3,
			0,
			0,
			3,
			1,
			0,
			0,
		}
		for rowIdx := range expectedRows {
			if res[rowIdx] != expectedRows[rowIdx] {
				t.Errorf(`Expected row %d to be %d, got %d`, rowIdx, expectedRows[rowIdx], res[rowIdx])
			}
		}
	})

}

func TestPart1(t *testing.T) {
	t.Run("test block 2", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day14.Part1(reader)
		if res != 136 {
			t.Errorf(`Expected number of blocks to be 136, got %d`, res)
		}
	})
}
