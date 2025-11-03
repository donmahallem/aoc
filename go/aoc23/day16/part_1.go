package day16

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type CellType uint8

type Position aoc_utils.Point[int]

const (
	CellTypeEmpty      CellType = '.'
	CellTypeRightUp    CellType = '/'
	CellTypeLeftUp     CellType = '\\'
	CellTypeVertical   CellType = '|'
	CellTypeHorizontal CellType = '-'
)

var dirUp Position = Position{X: 0, Y: -1}
var dirDown Position = Position{X: 0, Y: 1}
var dirLeft Position = Position{X: -1, Y: 0}
var dirRight Position = Position{X: 1, Y: 0}

type Movement struct {
	Pos Position
	Dir Position
}

type MovementMemory []uint8

const (
	dirBitRight uint8 = 1 << iota
	dirBitLeft
	dirBitUp
	dirBitDown
)

type Field struct {
	Cells  []CellType
	Width  int
	Height int
}

func ParseInputPart1(r io.Reader) Field {
	scanner := bufio.NewScanner(r)

	field := Field{
		Cells:  make([]CellType, 0, 64),
		Width:  0,
		Height: 0,
	}

	for scanner.Scan() {
		line := scanner.Bytes()
		if field.Width == 0 {
			field.Width = len(line)
		}
		field.Height++
		for i := 0; i < len(line); i++ {
			field.Cells = append(field.Cells, CellType(line[i]))
		}
	}
	return field
}

func dirMask(dir Position) uint8 {
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

func (f Field) index(pos Position) int {
	return pos.Y*f.Width + pos.X
}

func (f Field) cellAt(pos Position) *CellType {
	return &f.Cells[f.index(pos)]
}

func Simulate(field Field, memory MovementMemory, initial Movement) {
	stack := make([]Movement, 0, 64)
	stack = append(stack, initial)

	for len(stack) > 0 {
		mv := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		pos := mv.Pos
		dir := mv.Dir

		if pos.X < 0 || pos.X >= field.Width || pos.Y < 0 || pos.Y >= field.Height {
			continue
		}

		mask := dirMask(dir)
		if mask == 0 {
			continue
		}

		idx := field.index(pos)
		if memory[idx]&mask != 0 {
			continue
		}
		memory[idx] |= mask

		currentCell := field.cellAt(pos)

		switch *currentCell {
		case CellTypeHorizontal:
			if dir.Y == 0 {
				next := Position{X: pos.X + dir.X, Y: pos.Y + dir.Y}
				stack = append(stack, Movement{Pos: next, Dir: dir})
			} else {
				stack = append(stack,
					Movement{Pos: Position{X: pos.X - 1, Y: pos.Y}, Dir: dirLeft},
					Movement{Pos: Position{X: pos.X + 1, Y: pos.Y}, Dir: dirRight},
				)
			}
		case CellTypeVertical:
			if dir.X == 0 {
				next := Position{X: pos.X + dir.X, Y: pos.Y + dir.Y}
				stack = append(stack, Movement{Pos: next, Dir: dir})
			} else {
				stack = append(stack,
					Movement{Pos: Position{X: pos.X, Y: pos.Y - 1}, Dir: dirUp},
					Movement{Pos: Position{X: pos.X, Y: pos.Y + 1}, Dir: dirDown},
				)
			}
		case CellTypeRightUp:
			var turn Position
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
			next := Position{X: pos.X + turn.X, Y: pos.Y + turn.Y}
			stack = append(stack, Movement{Pos: next, Dir: turn})
		case CellTypeLeftUp:
			var turn Position
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
			next := Position{X: pos.X + turn.X, Y: pos.Y + turn.Y}
			stack = append(stack, Movement{Pos: next, Dir: turn})
		default:
			next := Position{X: pos.X + dir.X, Y: pos.Y + dir.Y}
			stack = append(stack, Movement{Pos: next, Dir: dir})
		}
	}
}

func Part1(in io.Reader) int {
	start := ParseInputPart1(in)
	memory := make(MovementMemory, start.Width*start.Height)
	Simulate(start, memory, Movement{Pos: Position{X: 0, Y: 0}, Dir: dirRight})
	return countEnergized(memory)
}

func countEnergized(memory MovementMemory) int {
	count := 0
	for _, mask := range memory {
		if mask != 0 {
			count++
		}
	}
	return count
}
