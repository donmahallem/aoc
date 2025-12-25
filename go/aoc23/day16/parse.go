package day16

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type cellType uint8

type position aoc_utils.Point[int]

const (
	cellTypeEmpty      cellType = '.'
	cellTypeRightUp    cellType = '/'
	cellTypeLeftUp     cellType = '\\'
	cellTypeVertical   cellType = '|'
	cellTypeHorizontal cellType = '-'
)

var dirUp position = position{X: 0, Y: -1}
var dirDown position = position{X: 0, Y: 1}
var dirLeft position = position{X: -1, Y: 0}
var dirRight position = position{X: 1, Y: 0}

type movement struct {
	Pos position
	Dir position
}

type movementMemory []uint8

const (
	dirBitRight uint8 = 1 << iota
	dirBitLeft
	dirBitUp
	dirBitDown
)

type field struct {
	Cells  []cellType
	Width  int
	Height int
}

func parseInput(r io.Reader) (*field, error) {
	scanner := bufio.NewScanner(r)

	f := field{
		Cells:  make([]cellType, 0, 64),
		Width:  0,
		Height: 0,
	}

	for scanner.Scan() {
		line := scanner.Bytes()
		if f.Width == 0 {
			f.Width = len(line)
		} else if f.Width != len(line) {
			return nil, aoc_utils.NewParseError("inconsistent line lengths in input", nil)
		}
		f.Height++
		for i := 0; i < len(line); i++ {
			f.Cells = append(f.Cells, cellType(line[i]))
		}
	}
	return &f, nil
}

func dirMask(dir position) uint8 {
	switch {
	case dir.X == 1 && dir.Y == 0:
		return dirBitRight
	case dir.X == -1 && dir.Y == 0:
		return dirBitLeft
	case dir.X == 0 && dir.Y == -1:
		return dirBitUp
	case dir.X == 0 && dir.Y == 1:
		return dirBitDown
	default:
		return 0
	}
}

func (f field) index(pos position) int {
	return pos.Y*f.Width + pos.X
}

func (f field) cellAt(pos position) *cellType {
	idx := f.index(pos)
	if idx < 0 || idx >= len(f.Cells) {
		return nil
	}
	return &f.Cells[idx]
}

func simulate(f field, memory movementMemory, initial movement) {
	stack := make([]movement, 0, 64)
	stack = append(stack, initial)

	for len(stack) > 0 {
		mv := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		pos := mv.Pos
		dir := mv.Dir

		if pos.X < 0 || pos.X >= f.Width || pos.Y < 0 || pos.Y >= f.Height {
			continue
		}

		mask := dirMask(dir)
		if mask == 0 {
			continue
		}

		idx := f.index(pos)
		if memory[idx]&mask != 0 {
			continue
		}
		memory[idx] |= mask

		currentCell := f.cellAt(pos)
		if currentCell == nil {
			continue
		}

		switch *currentCell {
		case cellTypeHorizontal:
			if dir.Y == 0 {
				next := position{X: pos.X + dir.X, Y: pos.Y + dir.Y}
				stack = append(stack, movement{Pos: next, Dir: dir})
			} else {
				stack = append(stack,
					movement{Pos: position{X: pos.X - 1, Y: pos.Y}, Dir: dirLeft},
					movement{Pos: position{X: pos.X + 1, Y: pos.Y}, Dir: dirRight},
				)
			}
		case cellTypeVertical:
			if dir.X == 0 {
				next := position{X: pos.X + dir.X, Y: pos.Y + dir.Y}
				stack = append(stack, movement{Pos: next, Dir: dir})
			} else {
				stack = append(stack,
					movement{Pos: position{X: pos.X, Y: pos.Y - 1}, Dir: dirUp},
					movement{Pos: position{X: pos.X, Y: pos.Y + 1}, Dir: dirDown},
				)
			}
		case cellTypeRightUp:
			var turn position
			switch {
			case dir.X == 1 && dir.Y == 0:
				turn = dirUp
			case dir.X == -1 && dir.Y == 0:
				turn = dirDown
			case dir.X == 0 && dir.Y == -1:
				turn = dirRight
			case dir.X == 0 && dir.Y == 1:
				turn = dirLeft
			default:
				continue
			}
			next := position{X: pos.X + turn.X, Y: pos.Y + turn.Y}
			stack = append(stack, movement{Pos: next, Dir: turn})
		case cellTypeLeftUp:
			var turn position
			switch {
			case dir.X == 1 && dir.Y == 0:
				turn = dirDown
			case dir.X == -1 && dir.Y == 0:
				turn = dirUp
			case dir.X == 0 && dir.Y == -1:
				turn = dirLeft
			case dir.X == 0 && dir.Y == 1:
				turn = dirRight
			default:
				continue
			}
			next := position{X: pos.X + turn.X, Y: pos.Y + turn.Y}
			stack = append(stack, movement{Pos: next, Dir: turn})
		default:
			next := position{X: pos.X + dir.X, Y: pos.Y + dir.Y}
			stack = append(stack, movement{Pos: next, Dir: dir})
		}
	}
}

func countEnergized(memory movementMemory) int {
	count := 0
	for _, mask := range memory {
		if mask != 0 {
			count++
		}
	}
	return count
}
