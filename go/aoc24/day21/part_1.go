package day21

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

var (
	NUMERIC_0         = Point{X: 1, Y: 3}
	NUMERIC_1         = Point{X: 0, Y: 2}
	NUMERIC_2         = Point{X: 1, Y: 2}
	NUMERIC_3         = Point{X: 2, Y: 2}
	NUMERIC_4         = Point{X: 0, Y: 1}
	NUMERIC_5         = Point{X: 1, Y: 1}
	NUMERIC_6         = Point{X: 2, Y: 1}
	NUMERIC_7         = Point{X: 0, Y: 0}
	NUMERIC_8         = Point{X: 1, Y: 0}
	NUMERIC_9         = Point{X: 2, Y: 0}
	NUMERIC_A         = Point{X: 2, Y: 3}
	DIRECTIONAL_LEFT  = Point{X: 0, Y: 1}
	DIRECTIONAL_RIGHT = Point{X: 2, Y: 1}
	DIRECTIONAL_UP    = Point{X: 1, Y: 0}
	DIRECTIONAL_DOWN  = Point{X: 1, Y: 1}
	DIRECTIONAL_A     = Point{X: 2, Y: 0}
)

type Point = aoc_utils.Point[int8]

func ParseInput(in io.Reader) *[][]Point {
	data := make([][]Point, 0)
	s := bufio.NewScanner(in)
	for s.Scan() {
		line := s.Bytes()
		data = append(data, make([]Point, 0, len(line)))
		for x := range len(line) {
			data[len(data)-1] = append(data[len(data)-1], *translateChar(&line[x]))
		}
	}
	return &data
}

func translateChar(b *byte) *Point {
	switch *b {
	case '0':
		return &NUMERIC_0
	case '1':
		return &NUMERIC_1
	case '2':
		return &NUMERIC_2
	case '3':
		return &NUMERIC_3
	case '4':
		return &NUMERIC_4
	case '5':
		return &NUMERIC_5
	case '6':
		return &NUMERIC_6
	case '7':
		return &NUMERIC_7
	case '8':
		return &NUMERIC_8
	case '9':
		return &NUMERIC_9
	case 'A':
		return &NUMERIC_A
	default:
		panic("Unexpected character provided")
	}
}

func WalkXFirst(start *Point, end *Point, dir *Point, currentDepth uint, maxDepth uint) (*Point, uint) {
	var steps, tempSteps uint
	steps = 0
	current := start
	var next *Point
	if dir.X != 0 {
		if dir.X < 0 {
			next = &DIRECTIONAL_LEFT
		} else {
			next = &DIRECTIONAL_RIGHT
		}
		for range aoc_utils.Abs(dir.X) {
			current, tempSteps = Walk(current, next, currentDepth+1, maxDepth)
			steps += tempSteps
		}
	}
	if dir.Y == 0 {
		// Fast exit
		// No Y moves
		current, tempSteps = Walk(current, &DIRECTIONAL_A, currentDepth+1, maxDepth)
		return current, steps + tempSteps
	} else if dir.Y < 0 {
		next = &DIRECTIONAL_UP
	} else {
		next = &DIRECTIONAL_DOWN
	}
	for range aoc_utils.Abs(dir.Y) {
		current, tempSteps = Walk(current, next, currentDepth+1, maxDepth)
		current = next
	}
	current, tempSteps = Walk(next, &DIRECTIONAL_A, currentDepth+1, maxDepth)
	steps += tempSteps
	return current, steps
}
func WalkYFirst(start *Point, end *Point, dir *Point, currentDepth uint, maxDepth uint) (*Point, uint) {
	var steps, tempSteps uint
	steps = 0
	current := start
	var next *Point
	if dir.Y != 0 {
		if dir.Y < 0 {
			next = &DIRECTIONAL_UP
		} else {
			next = &DIRECTIONAL_DOWN
		}
		for range aoc_utils.Abs(dir.Y) {
			current, tempSteps = Walk(current, next, currentDepth+1, maxDepth)
			steps += tempSteps
		}
	}
	if dir.X == 0 {
		// Fast exit
		// No Y moves
		current, tempSteps = Walk(current, &DIRECTIONAL_A, currentDepth+1, maxDepth)
		return current, steps + tempSteps
	} else if dir.X < 0 {
		next = &DIRECTIONAL_LEFT
	} else {
		next = &DIRECTIONAL_RIGHT
	}
	for range aoc_utils.Abs(dir.X) {
		current, tempSteps = Walk(current, next, currentDepth+1, maxDepth)
		steps += tempSteps
	}
	current, tempSteps = Walk(next, &DIRECTIONAL_A, currentDepth+1, maxDepth)
	steps += tempSteps
	return current, steps
}
func WalkNumeric(start *Point, end *Point, currentDepth uint, maxDepth uint) (*Point, uint) {
	dir := start.Diff(*end)
	if dir.X == 0 && dir.Y == 0 {
		// No walk necessary
		return Walk(&DIRECTIONAL_A, &DIRECTIONAL_A, currentDepth+1, maxDepth)
	}
	if start.X == 0 && end.Y == 3 {
		// First walk X than Y
		return WalkXFirst(start, end, dir, currentDepth, maxDepth)
	} else if end.X == 0 && start.Y == 3 {
		// First walk Y than X
		return WalkYFirst(start, end, dir, currentDepth, maxDepth)
	} else {
		// try first x and first y
		endSignA, endResultA := WalkXFirst(start, end, dir, currentDepth, maxDepth)
		if endSignB, endResultB := WalkYFirst(start, end, dir, currentDepth, maxDepth); endResultB < endResultA {
			return endSignB, endResultB
		}
		return endSignA, endResultA
	}
}

func WalkDirectional(start *Point, end *Point, currentDepth uint, maxDepth uint) (*Point, uint) {
	dir := start.Diff(*end)
	if dir.X == 0 && dir.Y == 0 {
		// No walk necessary
		return Walk(&DIRECTIONAL_A, &DIRECTIONAL_A, currentDepth+1, maxDepth)
	}
	if start.X == 0 && end.Y == 0 {
		// First walk X than Y
		return WalkXFirst(start, end, dir, currentDepth, maxDepth)
	} else if end.Y == 0 && start.X == 0 {
		// First walk Y than X
		return WalkYFirst(start, end, dir, currentDepth, maxDepth)
	} else {
		// try first x and first y
		endSignA, endResultA := WalkXFirst(start, end, dir, currentDepth, maxDepth)
		if endSignB, endResultB := WalkYFirst(start, end, dir, currentDepth, maxDepth); endResultB < endResultA {
			return endSignB, endResultB
		}
		return endSignA, endResultA
	}
}

func Walk(start *Point, end *Point, currentDepth uint, maxDepth uint) (*Point, uint) {
	if currentDepth == maxDepth {
		return start, 1
	} else if currentDepth == 0 {
		return WalkNumeric(start, end, currentDepth, maxDepth)
	} else {
		return WalkDirectional(start, end, currentDepth, maxDepth)
	}
}
