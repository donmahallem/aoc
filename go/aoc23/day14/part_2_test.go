package day14_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day14"
)

func TestParseInput2(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 10
		reader := strings.NewReader(testData)
		if res := day14.ParseInputPart1(reader); len(res) != expected {
			t.Errorf(`Expected %d to be %d`, len(res), expected)
		}

	})
	t.Run("test block 1", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day14.ParseInputPart1(reader)
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
		field := day14.ParseInputPart2(reader)
		day14.ApplyGravity(field, day14.DirUp)
		expectedField := day14.ParseInputPart2(strings.NewReader(testDataDirUp))
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
		field := day14.ParseInputPart2(reader)
		day14.ApplyGravity(field, day14.DirRight)
		expectedField := day14.ParseInputPart2(strings.NewReader(testDataDirRight))
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
		field := day14.ParseInputPart2(reader)
		day14.ApplyGravity(field, day14.DirDown)
		expectedField := day14.ParseInputPart2(strings.NewReader(testDataDirDown))
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
		field := day14.ParseInputPart2(reader)
		day14.ApplyGravity(field, day14.DirLeft)
		expectedField := day14.ParseInputPart2(strings.NewReader(testDataDirLeft))
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
		field := day14.ParseInputPart2(reader)
		day14.CycleDirections(field)
		expectedField := day14.ParseInputPart2(strings.NewReader(testDataCycle1))
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
		field := day14.ParseInputPart2(reader)
		day14.CycleDirections(field)
		day14.CycleDirections(field)
		expectedField := day14.ParseInputPart2(strings.NewReader(testDataCycle2))
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
		field := day14.ParseInputPart2(reader)
		day14.ApplyGravity(field, day14.DirUp)
		beamLoad := day14.CalculateBeamLoad(field)
		if beamLoad != 136 {
			t.Errorf(`Expected beam load to be 136, got %d`, beamLoad)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("test block 2", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day14.Part2(reader)
		if res != 64 {
			t.Errorf(`Expected number of blocks to be 64, got %d`, res)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		day14.Part2(reader)
		reader.Seek(0, 0)
	}
}
