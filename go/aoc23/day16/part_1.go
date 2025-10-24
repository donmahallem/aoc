package day16

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
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

type Cell struct {
	Type  CellType
	North *Cell
	East  *Cell
	South *Cell
	West  *Cell
}

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
	Cells  []Cell
	Width  int
	Height int
}

func ParseInputPart1(r io.Reader) Field {
	scanner := bufio.NewScanner(r)

	field := Field{
		Cells:  make([]Cell, 0, 64),
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
			field.Cells = append(field.Cells, Cell{Type: CellType(line[i])})
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

func (f Field) cellAt(pos Position) *Cell {
	return &f.Cells[f.index(pos)]
}

func SimulateStep(field Field, position Position, direction Position, memory MovementMemory) {
	if position.X < 0 || position.X >= field.Width || position.Y < 0 || position.Y >= field.Height {
		return
	}
	mask := dirMask(direction)
	if mask == 0 {
		return
	}
	idx := field.index(position)
	if memory[idx]&mask != 0 {
		return
	}
	memory[idx] |= mask

	currentCell := field.cellAt(position)
	switch currentCell.Type {
	case CellTypeHorizontal:
		if direction.Y == 0 {
			// pass through
			nextPosition := Position{X: position.X + direction.X, Y: position.Y + direction.Y}
			SimulateStep(field, nextPosition, direction, memory)
			return
		} else {
			// split beam
			SimulateStep(field, Position{X: position.X - 1, Y: position.Y}, dirLeft, memory)
			SimulateStep(field, Position{X: position.X + 1, Y: position.Y}, dirRight, memory)
			return
		}
	case CellTypeVertical:
		if direction.X == 0 {
			// pass through
			nextPosition := Position{X: position.X + direction.X, Y: position.Y + direction.Y}
			SimulateStep(field, nextPosition, direction, memory)
			return
		} else {
			// split beam
			SimulateStep(field, Position{X: position.X, Y: position.Y - 1}, dirUp, memory)
			SimulateStep(field, Position{X: position.X, Y: position.Y + 1}, dirDown, memory)
			return
		}
	case CellTypeRightUp:
		if direction.X > 0 {
			// right to up
			SimulateStep(field, Position{X: position.X, Y: position.Y - 1}, dirUp, memory)
			return
		} else if direction.Y < 0 {
			// up to right
			SimulateStep(field, Position{X: position.X + 1, Y: position.Y}, dirRight, memory)
			return
		} else if direction.X < 0 {
			// left to down
			SimulateStep(field, Position{X: position.X, Y: position.Y + 1}, dirDown, memory)
			return
		} else if direction.Y > 0 {
			// down to left
			SimulateStep(field, Position{X: position.X - 1, Y: position.Y}, dirLeft, memory)
			return
		}
	case CellTypeLeftUp:
		if direction.X > 0 {
			// right to down
			SimulateStep(field, Position{X: position.X, Y: position.Y + 1}, dirDown, memory)
			return
		} else if direction.Y < 0 {
			// up to left
			SimulateStep(field, Position{X: position.X - 1, Y: position.Y}, dirLeft, memory)
			return
		} else if direction.X < 0 {
			// left to up
			SimulateStep(field, Position{X: position.X, Y: position.Y - 1}, dirUp, memory)
			return
		} else if direction.Y > 0 {
			// down to right
			SimulateStep(field, Position{X: position.X + 1, Y: position.Y}, dirRight, memory)
			return
		}
	case CellTypeEmpty:
		// go ahead
		nextPosition := Position{X: position.X + direction.X, Y: position.Y + direction.Y}
		SimulateStep(field, nextPosition, direction, memory)
		return
	}
}

func Simulate(field Field) {
	currentPosition := Position{X: 0, Y: 0}
	currentDirection := Position{X: 1, Y: 0} // moving right

	for {
		if currentPosition.X < 0 || currentPosition.X >= field.Width || currentPosition.Y < 0 || currentPosition.Y >= field.Height {
			return
		}
		currentCell := field.cellAt(currentPosition)
		//nextPosition := Position{X: currentPosition.X + currentDirection.X, Y: currentPosition.Y + currentDirection.Y}
		switch currentCell.Type {
		case CellTypeHorizontal:
			if currentDirection.Y == 0 {
				// pass through
				currentPosition.X += currentDirection.X
				continue
			} else {
				// split beam
			}
		}
	}
}

func Part1(in io.Reader) int {
	start := ParseInputPart1(in)
	memory := make(MovementMemory, start.Width*start.Height)
	SimulateStep(start, Position{X: 0, Y: 0}, dirRight, memory)
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
