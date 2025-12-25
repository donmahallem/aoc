package day14

import (
	"strings"
	"testing"
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
		res, err := parseInputPart1(reader)
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
		if len(res) != expected {
			t.Errorf(`Expected %d to be %d`, len(res), expected)
		}

	})
	t.Run("test block 1", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res, err := parseInputPart1(reader)
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
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
