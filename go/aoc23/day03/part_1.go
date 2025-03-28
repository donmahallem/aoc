package day03

import (
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

type Match struct {
	value      int
	Start, End aoc_utils.Point[uint16]
	pair       bool
}

func IsNumber(b byte) bool {
	return b >= '0' && b <= '9'
}
func IsPart(b byte) bool {
	return !IsNumber(b) && b != '.'
}

func CheckForPair(field *aoc_utils.ByteField, match *Match) bool {
	var startX, endX uint16 = 0, (*field).Width - 1
	if (*match).Start.X > 0 {
		startX = (*match).Start.X - 1
	}
	if (*match).End.X < (*field).Width-1 {
		endX = (*match).End.X + 1
	}
	var y uint16 = (*match).Start.Y
	if (*match).Start.X > 0 && IsPart((*field).Field[y][startX]) {
		return true
	} else if (*match).Start.X < (*field).Width-1 && IsPart((*field).Field[y][endX]) {
		return true
	}
	for x := startX; x <= endX; x++ {
		if (*match).Start.Y > 1 {
			y = (*match).Start.Y - 1
			//fmt.Printf("Check Cell %d,%d\n", x, y)
			if IsPart((*field).Field[y][x]) {
				return true
			}
		}
		if (*match).Start.Y < (*field).Height-1 {
			y = (*match).Start.Y + 1
			//fmt.Printf("Check Cell %d,%d\n", x, y)
			if IsPart((*field).Field[y][x]) {
				return true
			}
		}
	}
	return false
}

func ParseField(in io.Reader) []Match {
	field, _ := aoc_utils.LoadField(in)
	matches := make([]Match, 0, 10)
	var currentMatch Match = Match{}
	currentMatch.value = -1
	for y := range field.Height {
		for x := range field.Width {
			if field.Field[y][x] >= '0' && field.Field[y][x] <= '9' {
				if currentMatch.value < 0 {
					currentMatch.value = int(field.Field[y][x] - '0')
					currentMatch.Start.X = x
					currentMatch.Start.Y = y
					currentMatch.End.Y = y
				} else {
					currentMatch.value = (currentMatch.value * 10) + int(field.Field[y][x]-'0')
					currentMatch.End.X = x
				}
			} else if currentMatch.value >= 0 {
				currentMatch.pair = CheckForPair(field, &currentMatch)
				matches = append(matches, currentMatch)
				currentMatch = Match{}
				currentMatch.value = -1
			}
		}
	}
	return matches
}

func Part1(in io.Reader) int {
	matches := ParseField(in)
	result := 0
	for _, match := range matches {
		if match.pair {
			result += match.value
		}
	}
	return result
}
