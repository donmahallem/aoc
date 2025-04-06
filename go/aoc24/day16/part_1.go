package day16

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

const (
	CELL_EMPTY uint8 = 0
	CELL_WALL  uint8 = 1
)

type Direction = aoc_utils.Point[int16]
type Field = [][]uint8
type Point = aoc_utils.Point[int16]
type Check struct {
	point Point
	dir   Direction
	score uint
}

var DIR_UP Direction = Direction{X: 0, Y: -1}
var DIR_DOWN Direction = Direction{X: 0, Y: 1}
var DIR_LEFT Direction = Direction{X: -1, Y: 0}
var DIR_RIGHT Direction = Direction{X: 1, Y: 0}

func ParseInput(in io.Reader) (Field, Point, Point) {
	field := make(Field, 0)
	s := bufio.NewScanner(in)
	start := Point{}
	end := Point{}
	for s.Scan() {
		line := s.Bytes()
		row := make([]uint8, len(line))
		for x := range len(line) {
			switch line[x] {
			case '#':
				row[x] = CELL_WALL
			case 'S':
				start.X = int16(x)
				start.Y = int16(len(field))
			case 'E':
				end.X = int16(x)
				end.Y = int16(len(field))
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
	case DIR_LEFT:
		return &DIR_DOWN
	case DIR_DOWN:
		return &DIR_RIGHT
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

func FindShortestPath(field *Field, start *Point, end *Point) uint {
	toCheck := make([]Check, 0)
	mapValues := make(map[Point]uint)
	toCheck = append(toCheck, Check{point: *start, dir: DIR_RIGHT, score: 0})
	nextCoord := Point{}
	checkDirs := make([]*Direction, 0, 3)
	for len(toCheck) > 0 {
		check := toCheck[0]
		toCheck = toCheck[1:]
		checkDirs = checkDirs[:0]
		checkDirs = append(checkDirs, &check.dir, translateLeft(&check.dir), translateRight(&check.dir))
		for checkDirIdx, checkDir := range checkDirs {
			nextCoord.X = check.point.X + checkDir.X
			nextCoord.Y = check.point.Y + checkDir.Y
			nextValue := check.score + 1
			if (*field)[nextCoord.Y][nextCoord.X] == CELL_WALL {
				continue
			}
			if checkDirIdx > 0 {
				nextValue += 1000
			}
			val, ok := mapValues[nextCoord]
			if ok && val < nextValue {
				continue
			}
			mapValues[nextCoord] = nextValue
			toCheck = append(toCheck, Check{dir: *checkDir, point: nextCoord, score: nextValue})
		}
	}
	return mapValues[*end]
}

func Part1(in io.Reader) int {
	field, start, end := ParseInput(in)
	return int(FindShortestPath(&field, &start, &end))
}
