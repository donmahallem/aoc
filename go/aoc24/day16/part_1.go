package day16

import (
	"bufio"
	"fmt"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

const (
	CELL_EMPTY int = -16
	CELL_WALL  int = -17
)

type direction = aoc_utils.Point[int16]
type field = [][]int
type point = aoc_utils.Point[int16]
type Check struct {
	point point
	dir   direction
	score int
}

var dirUP direction = direction{X: 0, Y: -1}
var dirDOWN direction = direction{X: 0, Y: 1}
var dirLEFT direction = direction{X: -1, Y: 0}
var dirRIGHT direction = direction{X: 1, Y: 0}
var dirsALL []direction = []direction{dirDOWN, dirLEFT, dirRIGHT, dirUP}

func printField(field *field, current *point, dir *direction) {
	for y := range int16(len(*field)) {
		for x := range int16(len((*field)[y])) {
			if current.X == x && current.Y == y {
				switch *dir {
				case dirDOWN:
					fmt.Print("v")
				case dirUP:
					fmt.Print("^")
				case dirRIGHT:
					fmt.Print(">")
				case dirLEFT:
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

type inputData struct {
	Field      field
	Start, End point
}

func parseInput(in io.Reader) (*inputData, error) {
	field := make(field, 0)
	s := bufio.NewScanner(in)
	start := point{}
	end := point{}
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
			case '.':
				row[x] = CELL_EMPTY
			default:
				return nil, aoc_utils.NewUnexpectedInputError(line[x])
			}
		}
		field = append(field, row)

	}
	return &inputData{Field: field, Start: start, End: end}, nil
}

func translateLeft(dir *direction) *direction {
	switch *dir {
	case dirUP:
		return &dirLEFT
	case dirRIGHT:
		return &dirUP
	case dirDOWN:
		return &dirRIGHT
	case dirLEFT:
		return &dirDOWN
	default:
		panic("Unknown Direction")
	}
}
func translateRight(dir *direction) *direction {
	switch *dir {
	case dirDOWN:
		return &dirLEFT
	case dirLEFT:
		return &dirUP
	case dirRIGHT:
		return &dirDOWN
	case dirUP:
		return &dirRIGHT
	default:
		panic("Unknown Direction")
	}
}

// Takes the 2D field and sets the path cost values on empty cells, that are
// reachable from the start
func calculatePathValues(field *field, start *point) {
	// Defensive checks to avoid panics on fuzzed inputs.
	if field == nil || len(*field) == 0 {
		return
	}
	if int(start.Y) < 0 || int(start.Y) >= len(*field) || int(start.X) < 0 || int(start.X) >= len((*field)[start.Y]) {
		return
	}
	toCheck := make([]Check, 0)
	(*field)[start.Y][start.X] = 0
	toCheck = append(toCheck, Check{point: *start, dir: dirRIGHT, score: 0})
	nextCoord := point{}
	checkDirs := make([]*direction, 0, 3)
	var currentFieldVal int
	// Walks iterativly breath first... was MUCH faster than depth first
	for len(toCheck) > 0 {
		check := toCheck[0]
		toCheck = toCheck[1:]
		checkDirs = checkDirs[:0]
		checkDirs = append(checkDirs,
			&check.dir,
			translateLeft(&check.dir),
			translateRight(&check.dir))
		for checkDirIdx, checkDir := range checkDirs {
			nextCoord.X = check.point.X + checkDir.X
			nextCoord.Y = check.point.Y + checkDir.Y
			// bounds check
			if int(nextCoord.Y) < 0 || int(nextCoord.Y) >= len(*field) || int(nextCoord.X) < 0 || int(nextCoord.X) >= len((*field)[nextCoord.Y]) {
				continue
			}
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

func Part1(in io.Reader) (int, error) {
	input, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	calculatePathValues(&input.Field, &input.Start)
	// bounds check before returning
	if int(input.End.Y) < 0 || int(input.End.Y) >= len(input.Field) || int(input.End.X) < 0 || int(input.End.X) >= len(input.Field[input.End.Y]) {
		return 0, aoc_utils.NewParseError("end coordinate out of bounds", nil)
	}
	return input.Field[input.End.Y][input.End.X], nil
}
