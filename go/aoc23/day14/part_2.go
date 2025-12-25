package day14

import (
	"hash/fnv"
	"io"
)

func applyGravity(f field, dir direction) {
	// defensive: ignore empty fields
	if len(f) == 0 || len(f[0]) == 0 {
		return
	}
	switch dir {
	case dirDown:
		for colIdx := range f[0] {
			lastEmptyIdx := len(f) - 1
			for rowIdx := len(f) - 1; rowIdx >= 0; rowIdx-- {
				switch f[rowIdx][colIdx] {
				case cellStone:
					if lastEmptyIdx != rowIdx {
						f[lastEmptyIdx][colIdx] = cellStone
						f[rowIdx][colIdx] = cellEmpty
					}
					lastEmptyIdx--
				case cellWall:
					lastEmptyIdx = rowIdx - 1
				}
			}
		}
	case dirUp:
		for colIdx := range f[0] {
			lastEmptyIdx := 0
			for rowIdx := range len(f) {
				switch f[rowIdx][colIdx] {
				case cellStone:
					if lastEmptyIdx != rowIdx {
						f[lastEmptyIdx][colIdx] = cellStone
						f[rowIdx][colIdx] = cellEmpty
					}
					lastEmptyIdx++
				case cellWall:
					lastEmptyIdx = rowIdx + 1
				}

			}
		}
	case dirLeft:
		for rowIdx := range f {
			lastEmptyIdx := 0
			for colIdx := range len(f[0]) {
				switch f[rowIdx][colIdx] {
				case cellStone:
					if lastEmptyIdx != colIdx {
						f[rowIdx][lastEmptyIdx] = cellStone
						f[rowIdx][colIdx] = cellEmpty
					}
					lastEmptyIdx++
				case cellWall:
					lastEmptyIdx = colIdx + 1
				}
			}
		}
	case dirRight:
		for rowIdx := range f {
			lastEmptyIdx := len(f[0]) - 1
			for colIdx := len(f[0]) - 1; colIdx >= 0; colIdx-- {
				switch f[rowIdx][colIdx] {
				case cellStone:
					if lastEmptyIdx != colIdx {
						f[rowIdx][lastEmptyIdx] = cellStone
						f[rowIdx][colIdx] = cellEmpty
					}
					lastEmptyIdx--
				case cellWall:
					lastEmptyIdx = colIdx - 1
				}
			}
		}
	}
}

func cycleDirections(f field) {
	applyGravity(f, dirUp)
	applyGravity(f, dirLeft)
	applyGravity(f, dirDown)
	applyGravity(f, dirRight)
}

func calculateBeamLoad(f field) uint {
	var accum uint = 0
	for rowIdx := range f {
		for colIdx := range f[rowIdx] {
			if f[rowIdx][colIdx] == cellStone {
				accum += uint(len(f) - rowIdx)
			}
		}
	}
	return accum
}

func encodeField(f field) uint64 {
	fnvHasher := fnv.New64()
	for rowIdx := range f {
		fnvHasher.Write(f[rowIdx])
	}
	return fnvHasher.Sum64()
}

type cycleMemory = int

const totalCycles cycleMemory = 1000000000

func Part2(in io.Reader) (uint, error) {
	f, err := parseInputPart2(in)
	if err != nil {
		return 0, err
	}

	memoryField := make(map[uint64]cycleMemory)
	memoryScore := make(map[cycleMemory]uint)
	/* Simulate subcycles and detect loops to skip ahead*/
	for cycleIdx := cycleMemory(0); cycleIdx < totalCycles; cycleIdx++ {
		cycleDirections(f)
		encoded := encodeField(f)
		if lastCycle, found := memoryField[encoded]; found {
			skipCycles := totalCycles - cycleIdx - 1
			cycleLength := cycleIdx - lastCycle
			skips := skipCycles / cycleLength
			cycleIdx += skips * cycleLength
			remainingCycles := (totalCycles - cycleIdx - 1) % cycleLength
			finalCycle := lastCycle + remainingCycles
			return memoryScore[finalCycle], nil
		} else {
			memoryField[encoded] = cycleIdx
			memoryScore[cycleIdx] = calculateBeamLoad(f)
		}
	}
	return 0, nil
}
