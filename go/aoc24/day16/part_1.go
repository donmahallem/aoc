package day16

import (
	"bufio"
	"fmt"
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

const (
	CELL_EMPTY int = -16
	CELL_WALL  int = -17
)

type Direction = aoc_utils.Point[int16]
type Field = [][]int
type Point = aoc_utils.Point[int16]
type PathValueMap = map[Point]int
type Check struct {
	point Point
	dir   Direction
	score int
}

var DIR_UP Direction = Direction{X: 0, Y: -1}
var DIR_DOWN Direction = Direction{X: 0, Y: 1}
var DIR_LEFT Direction = Direction{X: -1, Y: 0}
var DIR_RIGHT Direction = Direction{X: 1, Y: 0}
var DIRS_ALL []Direction = []Direction{DIR_DOWN, DIR_LEFT, DIR_RIGHT, DIR_UP}

func printField(field *Field, current *Point, dir *Direction) {
	for y := range int16(len(*field)) {
		for x := range int16(len((*field)[y])) {
			if current.X == x && current.Y == y {
				switch *dir {
				case DIR_DOWN:
					fmt.Print("v")
				case DIR_UP:
					fmt.Print("^")
				case DIR_RIGHT:
					fmt.Print(">")
				case DIR_LEFT:
					fmt.Print("<")
				}
				continue
			}
			switch (*field)[y][x] {
			case CELL_EMPTY:
				fmt.Print(".")
			case CELL_WALL:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
func ParseInput(in io.Reader) (Field, Point, Point) {
	field := make(Field, 0)
	s := bufio.NewScanner(in)
	start := Point{}
	end := Point{}
	for s.Scan() {
		line := s.Bytes()
		row := make([]int, len(line))
		for x := range len(line) {
			switch line[x] {
			case '#':
				row[x] = CELL_WALL
			case 'S':
				start.X = int16(x)
				start.Y = int16(len(field))
				row[x] = CELL_EMPTY
			case 'E':
				end.X = int16(x)
				end.Y = int16(len(field))
				row[x] = CELL_EMPTY
			default:
				row[x] = CELL_EMPTY
			}
		}
		field = append(field, row)

	}
	return field, start, end
}

func translateLeft(dir *Direction) *Direction {
	switch *dir {
	case DIR_UP:
		return &DIR_LEFT
	case DIR_RIGHT:
		return &DIR_UP
	case DIR_DOWN:
		return &DIR_RIGHT
	case DIR_LEFT:
		return &DIR_DOWN
	default:
		panic("Unknown Direction")
	}
}
func translateRight(dir *Direction) *Direction {
	switch *dir {
	case DIR_DOWN:
		return &DIR_LEFT
	case DIR_LEFT:
		return &DIR_UP
	case DIR_RIGHT:
		return &DIR_DOWN
	case DIR_UP:
		return &DIR_RIGHT
	default:
		panic("Unknown Direction")
	}
}

func CalculatePathValues(field *Field, start *Point) {
	toCheck := make([]Check, 0)
	(*field)[start.Y][start.X] = 0
	toCheck = append(toCheck, Check{point: *start, dir: DIR_RIGHT, score: 0})
	nextCoord := Point{}
	checkDirs := make([]*Direction, 0, 3)
	var currentFieldVal int
	for len(toCheck) > 0 {
		check := toCheck[0]
		toCheck = toCheck[1:]
		checkDirs = checkDirs[:0]
		checkDirs = append(checkDirs, &check.dir, translateLeft(&check.dir), translateRight(&check.dir))
		for checkDirIdx, checkDir := range checkDirs {
			nextCoord.X = check.point.X + checkDir.X
			nextCoord.Y = check.point.Y + checkDir.Y
			currentFieldVal = (*field)[nextCoord.Y][nextCoord.X]
			if currentFieldVal == CELL_WALL {
				continue
			}
			nextValue := check.score + 1
			if checkDirIdx > 0 {
				nextValue += 1000
			}
			if currentFieldVal != CELL_EMPTY && currentFieldVal < nextValue {
				continue
			}
			(*field)[nextCoord.Y][nextCoord.X] = nextValue
			toCheck = append(toCheck, Check{dir: *checkDir, point: nextCoord, score: nextValue})
		}
	}
}

func Part1(in io.Reader) int {
	field, start, end := ParseInput(in)
	CalculatePathValues(&field, &start)

	return field[end.Y][end.X]
}
