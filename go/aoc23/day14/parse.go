package day14

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type field [][]cell
type direction uint8

const (
	dirLeft  direction = 'L'
	dirRight direction = 'R'
	dirDown  direction = 'D'
	dirUp    direction = 'U'
)

type cell = byte

const (
	cellEmpty cell = '.'
	cellStone cell = 'O'
	cellWall  cell = '#'
)

func parseInputPart1(r io.Reader) ([]uint, error) {
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
		if len(line) != len(lastEmpties) {
			return nil, aoc_utils.NewParseError("inconsistent line lengths in input", nil)
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
	return stonesPerRow, nil
}

func parseInputPart2(r io.Reader) (field, error) {
	scanner := bufio.NewScanner(r)

	var f field = make(field, 0, 16)
	width := -1
	for scanner.Scan() {
		row := append([]cell(nil), scanner.Bytes()...)
		if width == -1 {
			width = len(row)
		} else if len(row) != width {
			return nil, aoc_utils.NewParseError("line length is uneven", nil)
		}
		f = append(f, row)
	}
	return f, nil
}
