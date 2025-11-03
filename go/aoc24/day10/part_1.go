package day10

import (
	"bufio"
	"io"
	"reflect"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type Position aoc_utils.Point[uint16]

type Field struct {
	width, height uint16
	Field         [][]byte
	starts        []Position
}

func (a *Field) Compare(f *Field) bool {
	return a.width == f.width && a.height == f.height && reflect.DeepEqual(a.Field, f.Field) && reflect.DeepEqual(a.starts, f.starts)
}

func NewField(width uint16, height uint16, field [][]byte, starts []Position) Field {
	return Field{width: width, height: height, Field: field, starts: starts}
}

func LoadField(reader io.Reader) (Field, error) {
	obstacles := make([][]byte, 0)
	starts := make([]Position, 0)
	s := bufio.NewScanner(reader)
	y := uint16(0)
	width := uint16(0)
	for s.Scan() {
		line := s.Bytes()
		if width == 0 {
			width = uint16(len(line))
		} else if width != uint16(len(line)) {
			panic("Line length is uneven")
		}
		row := make([]byte, len(line))
		for idx, character := range line {
			row[idx] = character - '0'
			if row[idx] == 0 {
				starts = append(starts, Position{X: uint16(idx), Y: y})
			}
		}
		obstacles = append(obstacles, row)
		y++
	}
	return NewField(width, y, obstacles, starts), nil
}

func WalkDepth(data *Field, x uint16, y uint16, depth uint8, ends *map[Position]bool) int {
	if depth == 9 {
		(*ends)[Position{X: x, Y: y}] = true
		return 1
	}
	result := 0
	nextDepth := depth + 1
	if x > 0 && (*data).Field[y][x-1] == nextDepth {
		result += WalkDepth(data, x-1, y, nextDepth, ends)
	}
	if y > 0 && (*data).Field[y-1][x] == nextDepth {
		result += WalkDepth(data, x, y-1, nextDepth, ends)
	}
	if x < (*data).width-1 && (*data).Field[y][x+1] == nextDepth {
		result += WalkDepth(data, x+1, y, nextDepth, ends)
	}
	if y < (*data).height-1 && (*data).Field[y+1][x] == nextDepth {
		result += WalkDepth(data, x, y+1, nextDepth, ends)
	}
	return result
}

func SearchAll(field *Field) int {
	result := 0
	for i := range len((*field).starts) {
		mapper := make(map[Position]bool)
		WalkDepth(field, (*field).starts[i].X, (*field).starts[i].Y, 0, &mapper)
		result += len(mapper)
	}
	return result
}

func Part1(in io.Reader) int {
	data, _ := LoadField(in)
	return SearchAll(&data)
}
