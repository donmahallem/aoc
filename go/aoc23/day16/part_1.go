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

type MovementMemory map[Position]map[Position]bool

type Field struct {
	Cells  [][]Cell
	Width  int
	Height int
}

func ParseInputPart1(r io.Reader) Field {
	scanner := bufio.NewScanner(r)

	field := Field{
		Cells:  make([][]Cell, 0),
		Width:  0,
		Height: 0,
	}

	for scanner.Scan() {
		line := scanner.Bytes()
		if field.Width == 0 {
			field.Width = len(line)
		}
		field.Height++
		parsedLine := make([]Cell, len(line))
		for i := range line {
			parsedLine[i] = Cell{Type: CellType(line[i])}
		}
		field.Cells = append(field.Cells, parsedLine)
	}
	return field
}

func SimulateStep(field Field, position Position, direction Position, memory MovementMemory) {
	if position.X < 0 || position.X >= field.Width || position.Y < 0 || position.Y >= field.Height {
		return
	}
	if mem, ok := memory[position]; ok {
		if _, ok2 := memory[position][direction]; ok2 {
			return
		}
		mem[direction] = true
	} else {
		mem2 := make(map[Position]bool)
		mem2[direction] = true
		memory[position] = mem2

	}
	currentCell := &field.Cells[position.Y][position.X]
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
		currentCell := &field.Cells[currentPosition.Y][currentPosition.X]
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
	memory := make(MovementMemory)
	SimulateStep(start, Position{X: 0, Y: 0}, dirRight, memory)
	return len(memory)
}
