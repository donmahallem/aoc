package day14

import (
	"bufio"
	"io"
)

func ParseInput(r io.Reader) []uint {
	scanner := bufio.NewScanner(r)

	var lastEmpties []uint16 = nil
	stonesPerRow := make([]uint, 0, 16)

	var currentRow uint16 = 0
	for scanner.Scan() {
		stonesPerRow = append(stonesPerRow, 0)
		line := scanner.Bytes()
		if lastEmpties == nil {
			lastEmpties = make([]uint16, len(line))
		}
		for idx, c := range line {
			switch c {
			case 'O':
				stonesPerRow[lastEmpties[idx]] += 1
				lastEmpties[idx] += 1
			case '#':
				lastEmpties[idx] = uint16(currentRow) + 1
			}
		}
		currentRow++
	}
	return stonesPerRow
}

func Part1(in io.Reader) uint {
	start := ParseInput(in)
	accum := uint(0)
	for idx, val := range start {
		accum += val * uint(len(start)-idx)
	}
	return accum
}
