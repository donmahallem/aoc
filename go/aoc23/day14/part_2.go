package day14

import (
	"bufio"
	"hash/fnv"
	"io"
)

type Field [][]Cell
type Direction uint8

const (
	DirLeft  Direction = 'L'
	DirRight Direction = 'R'
	DirDown  Direction = 'D'
	DirUp    Direction = 'U'
)

type Cell = byte

const (
	CellEmpty Cell = '.'
	CellStone Cell = 'O'
	CellWall  Cell = '#'
)

func ApplyGravity(field Field, dir Direction) {

	if dir == DirDown {
		for colIdx := range field[0] {
			lastEmptyIdx := len(field) - 1
			for rowIdx := len(field) - 1; rowIdx >= 0; rowIdx-- {
				switch field[rowIdx][colIdx] {
				case CellStone:
					if lastEmptyIdx != rowIdx {
						field[lastEmptyIdx][colIdx] = CellStone
						field[rowIdx][colIdx] = CellEmpty
					}
					lastEmptyIdx--
				case CellWall:
					lastEmptyIdx = rowIdx - 1
				}
			}
		}
	} else if dir == DirUp {
		for colIdx := range field[0] {
			lastEmptyIdx := 0
			for rowIdx := 0; rowIdx < len(field); rowIdx++ {
				switch field[rowIdx][colIdx] {
				case CellStone:
					if lastEmptyIdx != rowIdx {
						field[lastEmptyIdx][colIdx] = CellStone
						field[rowIdx][colIdx] = CellEmpty
					}
					lastEmptyIdx++
				case CellWall:
					lastEmptyIdx = rowIdx + 1
				}

			}
		}
	} else if dir == DirLeft {
		for rowIdx := range field {
			lastEmptyIdx := 0
			for colIdx := 0; colIdx < len(field[0]); colIdx++ {
				switch field[rowIdx][colIdx] {
				case CellStone:
					if lastEmptyIdx != colIdx {
						field[rowIdx][lastEmptyIdx] = CellStone
						field[rowIdx][colIdx] = CellEmpty
					}
					lastEmptyIdx++
				case CellWall:
					lastEmptyIdx = colIdx + 1
				}
			}
		}
	} else if dir == DirRight {
		for rowIdx := range field {
			lastEmptyIdx := len(field[0]) - 1
			for colIdx := len(field[0]) - 1; colIdx >= 0; colIdx-- {
				switch field[rowIdx][colIdx] {
				case CellStone:
					if lastEmptyIdx != colIdx {
						field[rowIdx][lastEmptyIdx] = CellStone
						field[rowIdx][colIdx] = CellEmpty
					}
					lastEmptyIdx--
				case CellWall:
					lastEmptyIdx = colIdx - 1
				}
			}
		}
	}
}

func CycleDirections(field Field) {
	ApplyGravity(field, DirUp)
	ApplyGravity(field, DirLeft)
	ApplyGravity(field, DirDown)
	ApplyGravity(field, DirRight)
}

func ParseInputPart2(r io.Reader) Field {
	scanner := bufio.NewScanner(r)

	var field Field = make(Field, 0, 16)
	for scanner.Scan() {
		/**
		 * Create a copy of the scanned bytes to avoid modifying the underlying buffer
		 * on the next scan.
		 */
		row := append([]Cell(nil), scanner.Bytes()...)
		field = append(field, row)
	}
	return field
}

func CalculateBeamLoad(field Field) uint {
	var accum uint = 0
	for rowIdx := range field {
		for colIdx := range field[rowIdx] {
			if field[rowIdx][colIdx] == CellStone {
				accum += uint(len(field) - rowIdx)
			}
		}
	}
	return accum
}

func encodeField(field Field) uint64 {
	fnvHasher := fnv.New64()
	for rowIdx := range field {
		fnvHasher.Write(field[rowIdx])
	}
	return fnvHasher.Sum64()
}

type CycleMemory = int

const totalCycles CycleMemory = 1000000000

func Part2(in io.Reader) uint {
	field := ParseInputPart2(in)

	memory := make(map[uint64]CycleMemory)
	/* Simulate subcycles and detect loops to skip ahead*/
	for cycleIdx := CycleMemory(0); cycleIdx < totalCycles; cycleIdx++ {
		CycleDirections(field)
		encoded := encodeField(field)
		if lastCycle, found := memory[encoded]; found {
			skipCycles := totalCycles - cycleIdx - 1
			cycleLength := cycleIdx - lastCycle
			skips := skipCycles / cycleLength
			cycleIdx += skips * cycleLength
		} else {
			memory[encoded] = cycleIdx
		}
	}

	return CalculateBeamLoad(field)
}
