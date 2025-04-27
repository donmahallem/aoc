package day21

import "fmt"

var numericLookup map[byte]Point = map[byte]Point{
	'0': {X: 1, Y: 3},
	'1': {X: 0, Y: 3},
	'2': {X: 1, Y: 3},
	'3': {X: 2, Y: 1},
	'4': {X: 0, Y: 1},
	'5': {X: 1, Y: 1},
	'6': {X: 2, Y: 1},
	'7': {X: 0, Y: 0},
	'8': {X: 1, Y: 0},
	'9': {X: 2, Y: 0},
	'A': {X: 2, Y: 3},
}

var (
	NUMERIC_0 = Point{X: 1, Y: 3}
	NUMERIC_1 = Point{X: 0, Y: 2}
	NUMERIC_2 = Point{X: 1, Y: 2}
	NUMERIC_3 = Point{X: 2, Y: 2}
	NUMERIC_4 = Point{X: 0, Y: 1}
	NUMERIC_5 = Point{X: 1, Y: 1}
	NUMERIC_6 = Point{X: 2, Y: 1}
	NUMERIC_7 = Point{X: 0, Y: 0}
	NUMERIC_8 = Point{X: 1, Y: 0}
	NUMERIC_9 = Point{X: 2, Y: 0}
	NUMERIC_A = Point{X: 2, Y: 3}
)

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

func WalkNumericSequenceSub(start *Point, end *Point, currentDepth uint, maxDepth uint) uint {
	dir := start.Diff(*end)
	if dir.X == 0 && dir.Y == 0 {
		// No walking necessary
		return WalkDirectional(&DIRECTIONAL_A, &DIRECTIONAL_A, currentDepth+1, maxDepth)
	}
	if start.X == 0 && end.Y == 3 {
		// First walk X than Y
		return walkXFirstDirectional(start, end, dir, currentDepth, maxDepth)
	} else if end.X == 0 && start.Y == 3 {
		// First walk Y than X
		return walkYFirstDirectional(start, end, dir, currentDepth, maxDepth)
	} else {
		// try first x and first y
		endResultA := walkXFirstDirectional(start, end, dir, currentDepth, maxDepth)
		if endResultB := walkYFirstDirectional(start, end, dir, currentDepth, maxDepth); endResultB < endResultA {
			return endResultB
		}
		return endResultA
	}
}

func WalkNumericSequence(sequence *[]byte, currentDepth uint, maxDepth uint) uint {
	var count uint = 0
	var previousLetter byte = 'A'
	for _, key := range *sequence {
		fmt.Printf("Walk from %c to %c\n", previousLetter, key)
		from := numericLookup[previousLetter]
		to := numericLookup[key]
		count += WalkNumericSequenceSub(&from, &to, currentDepth, maxDepth)
		previousLetter = key
	}
	return count
}
