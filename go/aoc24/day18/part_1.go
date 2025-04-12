package day18

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/aoc_utils"
)

const CELL_CORRUPTED int = -1

var DIRS_ALL [4]Point = [4]Point{Point{X: -1, Y: 0}, Point{X: 0, Y: -1}, Point{X: 1, Y: 0}, Point{X: 0, Y: 1}}

type Field = [][]int
type Point = aoc_utils.Point[int]

func ParseInput(in io.Reader) *[]Point {
	points := make([]Point, 0)
	s := bufio.NewScanner(in)
	for s.Scan() {
		line := s.Text()
		commaIndex := strings.Index(line, ",")
		point := Point{}
		point.X, _ = strconv.Atoi(line[0:commaIndex])
		point.Y, _ = strconv.Atoi(line[commaIndex+1:])
		points = append(points, point)
	}
	return &points
}

func ConvertInputToField(points *[]Point, steps, width, height uint) *Field {
	field := make(Field, 0, height)
	for range height {
		field = append(field, make([]int, width))
	}
	for step := range steps {
		field[(*points)[step].Y][(*points)[step].X] = CELL_CORRUPTED
	}
	return &field
}

func FindShortestPath(field *Field) int {
	fieldHeight := len(*field)
	fieldWidth := len((*field)[0])
	toCheck := make([]Point, 0)
	toCheck = append(toCheck, Point{X: 0, Y: 0})
	var currentCoord, nextCoord Point
	var currentValue, nextValue int
	for len(toCheck) > 0 {
		currentCoord = toCheck[0]
		toCheck = toCheck[1:]
		currentValue = (*field)[currentCoord.Y][currentCoord.X]
		nextValue = currentValue + 1
		for _, checkDir := range DIRS_ALL {
			nextCoord.X = currentCoord.X + checkDir.X
			nextCoord.Y = currentCoord.Y + checkDir.Y
			if nextCoord.X < 0 || nextCoord.Y < 0 || nextCoord.X >= fieldWidth || nextCoord.Y >= fieldHeight {
				// next coord outside the field dimensions
				continue
			} else if currentNextValue := (*field)[nextCoord.Y][nextCoord.X]; currentNextValue == CELL_CORRUPTED ||
				(currentNextValue <= nextValue && currentNextValue > 0) {
				// CELL IS ALREADY CORUPTED OR A Lower value was found
				continue
			}
			(*field)[nextCoord.Y][nextCoord.X] = nextValue
			toCheck = append(toCheck, nextCoord)
		}
	}
	return (*field)[fieldHeight-1][fieldWidth-1]
}

func Part1(in io.Reader) int {
	points := ParseInput(in)
	field := ConvertInputToField(points, 1024, 71, 71)
	return FindShortestPath(field)
}
