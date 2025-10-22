package day13_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day13"
)

var testData string = `
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func TestParseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 2
		reader := strings.NewReader(testData)
		if res := day13.ParseInput(reader); len(res) != expected {
			t.Errorf(`Expected %d to be %d`, len(res), expected)
		}

	})
	t.Run("test block 1", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day13.ParseInput(reader)
		expectedRows := []int{
			0b11001101,
			0b10110100,
			0b100000011,
			0b100000011,
			0b010110100,
			0b011001100,
			0b010110101,
		}
		expectedColumns := []int{
			0b1001101, //0
			0b0001100, //1
			0b1110011, //2
			0b0100001, //3
			0b1010010, //4
			0b1010010, //5
			0b0100001, //6
			0b1110011, //7
			0b0001100, //8
		}
		for rowIdx := range expectedRows {
			if res[0].Rows[rowIdx] != expectedRows[rowIdx] {
				t.Errorf(`Expected row %d to be %b, got %b`, rowIdx, expectedRows[rowIdx], res[0].Rows[rowIdx])
			}
		}
		for colIdx := range expectedColumns {
			if res[0].Cols[colIdx] != expectedColumns[colIdx] {
				t.Errorf(`Expected row %d to be %b, got %b`, colIdx, expectedColumns[colIdx], res[0].Cols[colIdx])
			}
		}
	})

}

func TestPart1(t *testing.T) {
	t.Run("test block 2", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day13.Part1(reader)
		if res != 405 {
			t.Errorf(`Expected number of blocks to be 405, got %d`, res)
		}
	})
}
