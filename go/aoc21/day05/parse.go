package day05

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type coords = [5]int

type parserState uint8

const (
	stateInitial parserState = iota
	stateNumberX1
	stateNumberX2
	stateNumberY1
	stateNumberY2
)

func parseInput(in io.Reader) ([]coords, error) {
	numbers := make([]coords, 0, 128)
	// expected comma seperated list of numbers
	s := bufio.NewScanner(in)
	for s.Scan() {
		currentState := stateInitial
		line := s.Bytes()
		if len(line) == 0 {
			continue
		}
		var c coords
		for i := 0; i < len(line); i++ {
			switch currentState {
			case stateInitial:
				if line[i] >= '0' && line[i] <= '9' {
					currentState = stateNumberX1
					c[0] = int(line[i] - '0')
				} else {
					return nil, aoc_utils.NewParseErrorf("Unexpected character in initial state %c", line[i])
				}
			case stateNumberX1:
				if line[i] >= '0' && line[i] <= '9' {
					c[0] = (c[0] * 10) + int(line[i]-'0')
				} else if line[i] == ',' {
					currentState = stateNumberY1
				} else {
					return nil, aoc_utils.NewParseErrorf("Unexpected character in stateNumberX1 %c", line[i])
				}
			case stateNumberY1:
				if line[i] >= '0' && line[i] <= '9' {
					c[1] = (c[1] * 10) + int(line[i]-'0')
				} else if line[i] == ' ' {
					currentState = stateNumberX2
					i += 3
				} else {
					return nil, aoc_utils.NewParseErrorf("Unexpected character in stateNumberY1 %c", line[i])
				}
			case stateNumberX2:
				if line[i] >= '0' && line[i] <= '9' {
					c[2] = (c[2] * 10) + int(line[i]-'0')
				} else if line[i] == ',' {
					currentState = stateNumberY2
				} else {
					return nil, aoc_utils.NewParseErrorf("Unexpected character in stateNumberX2 %c", line[i])
				}
			case stateNumberY2:
				if line[i] >= '0' && line[i] <= '9' {
					c[3] = (c[3] * 10) + int(line[i]-'0')
				} else {
					return nil, aoc_utils.NewParseErrorf("Unexpected character in stateNumberY2 %c", line[i])
				}

			}

		}
		// normalize the arrows such that they are x+
		if c[0] > c[2] {
			c[0], c[1], c[2], c[3] = c[2], c[3], c[0], c[1]
		}
		getDirection(&c)
		numbers = append(numbers, c)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return numbers, nil
}

type Direction = int

const (
	DirectionHorizontal   Direction = 1
	DirectionVertical     Direction = 2
	DirectionDiagonalUp   Direction = 4
	DirectionDiagonalDown Direction = 8
	DirectionDiagonal     Direction = DirectionDiagonalUp | DirectionDiagonalDown
	DirectionOther        Direction = 16
)

func getDirection(arr *coords) {
	if arr[1] == arr[3] {
		arr[4] = DirectionHorizontal
	} else if arr[0] == arr[2] {
		arr[4] = DirectionVertical
	} else if arr[2]-arr[0] == arr[3]-arr[1] {
		arr[4] = DirectionDiagonalUp
	} else if arr[2]-arr[0] == -(arr[3] - arr[1]) {
		arr[4] = DirectionDiagonalDown
	} else {
		arr[4] = DirectionOther // default case, should not happen for valid input
	}
}

type point struct {
	x, y int
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func signInt(n int) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
}
func pointOnSegment(seg coords, px, py int) bool {
	minX, maxX := min(seg[0], seg[2]), max(seg[0], seg[2])
	minY, maxY := min(seg[1], seg[3]), max(seg[1], seg[3])
	if px < minX || px > maxX || py < minY || py > maxY {
		return false
	}
	// Collinearity check for axis-aligned or 45-degree segments:
	// (px - x1)*(y2 - y1) == (py - y1)*(x2 - x1)
	return (px-seg[0])*(seg[3]-seg[1]) == (py-seg[1])*(seg[2]-seg[0])
}
