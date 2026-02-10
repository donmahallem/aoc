package day14

import (
	"strings"
	"testing"
)

func TestParseInput2(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 10
		reader := strings.NewReader(testData)
		res, err := parseInputPart2(reader)
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

const testDataDirUp string = `OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`
const testDataDirRight string = `....O#....
.OOO#....#
.....##...
.OO#....OO
......OO#.
.O#...O#.#
....O#..OO
.........O
#....###..
#..OO#....`
const testDataDirDown string = `.....#....
....#....#
...O.##...
...#......
O.O....O#O
O.#..O.#.#
O....#....
OO....OO..
#OO..###..
#OO.O#...O`

const testDataDirLeft string = `O....#....
OOO.#....#
.....##...
OO.#OO....
OO......#.
O.#O...#.#
O....#OO..
O.........
#....###..
#OO..#....`

const testDataCycle1 string = `.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....`
const testDataCycle2 string = `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O`

func TestApplyGravity(t *testing.T) {
	t.Run("apply gravity up", func(t *testing.T) {
		reader := strings.NewReader(testData)
		field, err := parseInputPart2(reader)
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
		applyGravity(field, dirUp)
		expectedField, err := parseInputPart2(strings.NewReader(testDataDirUp))
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
		for rowIdx := range expectedField {
			for colIdx := range expectedField[rowIdx] {
				if field[rowIdx][colIdx] != expectedField[rowIdx][colIdx] {
					t.Errorf(`Expected cell (%d,%d) to be %c, got %c`, rowIdx, colIdx, expectedField[rowIdx][colIdx], field[rowIdx][colIdx])
				}
			}
		}
	})
	t.Run("apply gravity right", func(t *testing.T) {
		reader := strings.NewReader(testData)
		field, err := parseInputPart2(reader)
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
		applyGravity(field, dirRight)
		expectedField, err := parseInputPart2(strings.NewReader(testDataDirRight))
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
		for rowIdx := range expectedField {
			for colIdx := range expectedField[rowIdx] {
				if field[rowIdx][colIdx] != expectedField[rowIdx][colIdx] {
					t.Errorf(`Expected cell (%d,%d) to be %c, got %c`, rowIdx, colIdx, expectedField[rowIdx][colIdx], field[rowIdx][colIdx])
				}
			}
		}
	})
	t.Run("apply gravity down", func(t *testing.T) {
		reader := strings.NewReader(testData)
		field, err := parseInputPart2(reader)
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
		applyGravity(field, dirDown)
		expectedField, err := parseInputPart2(strings.NewReader(testDataDirDown))
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
		for rowIdx := range expectedField {
			for colIdx := range expectedField[rowIdx] {
				if field[rowIdx][colIdx] != expectedField[rowIdx][colIdx] {
					t.Errorf(`Expected cell (%d,%d) to be %c, got %c`, rowIdx, colIdx, expectedField[rowIdx][colIdx], field[rowIdx][colIdx])
				}
			}
		}
	})
	t.Run("apply gravity left", func(t *testing.T) {
		reader := strings.NewReader(testData)
		field, err := parseInputPart2(reader)
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
		applyGravity(field, dirLeft)
		expectedField, err := parseInputPart2(strings.NewReader(testDataDirLeft))
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
		for rowIdx := range expectedField {
			for colIdx := range expectedField[rowIdx] {
				if field[rowIdx][colIdx] != expectedField[rowIdx][colIdx] {
					t.Errorf(`Expected cell (%d,%d) to be %c, got %c`, rowIdx, colIdx, expectedField[rowIdx][colIdx], field[rowIdx][colIdx])
				}
			}
		}
	})
}
func TestCycleDirections(t *testing.T) {
	t.Run("test cycle directions once", func(t *testing.T) {
		reader := strings.NewReader(testData)
		field, err := parseInputPart2(reader)
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
		cycleDirections(field)
		expectedField, err := parseInputPart2(strings.NewReader(testDataCycle1))
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
		for rowIdx := range expectedField {
			for colIdx := range expectedField[rowIdx] {
				if field[rowIdx][colIdx] != expectedField[rowIdx][colIdx] {
					t.Errorf(`Expected cell (%d,%d) to be %c, got %c`, rowIdx, colIdx, expectedField[rowIdx][colIdx], field[rowIdx][colIdx])
				}
			}
		}
	})
	t.Run("test cycle directions twice", func(t *testing.T) {
		reader := strings.NewReader(testData)
		field, err := parseInputPart2(reader)
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
		cycleDirections(field)
		cycleDirections(field)
		expectedField, err := parseInputPart2(strings.NewReader(testDataCycle2))
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
		for rowIdx := range expectedField {
			for colIdx := range expectedField[rowIdx] {
				if field[rowIdx][colIdx] != expectedField[rowIdx][colIdx] {
					t.Errorf(`Expected cell (%d,%d) to be %c, got %c`, rowIdx, colIdx, expectedField[rowIdx][colIdx], field[rowIdx][colIdx])
				}
			}
		}
	})
}

func TestCalculateBeamLoad(t *testing.T) {
	t.Run("test beam load", func(t *testing.T) {
		reader := strings.NewReader(testData)
		field, err := parseInputPart2(reader)
		if err != nil {
			t.Fatalf("Unexpected parse error: %v", err)
		}
		applyGravity(field, dirUp)
		beamLoad := calculateBeamLoad(field)
		if beamLoad != 136 {
			t.Errorf(`Expected beam load to be 136, got %d`, beamLoad)
		}
	})
}
