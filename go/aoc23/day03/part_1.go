package day03

import (
	"io"
	"slices"

	"github.com/donmahallem/aoc/aoc_utils"
)

const (
	TYPE_NUM uint8 = iota
	TYPE_PART
	TYPE_SPACE
)

type Match struct {
	Value           int
	StartX, EndX, Y int16
	Pair            bool
}

func NewMatch(t int, startX, endX, y int16) *Match {
	return &Match{t, startX, endX, y, false}
}

func GetType(b byte) uint8 {
	if b == '.' {
		return TYPE_SPACE
	} else if b >= '0' && b <= '9' {
		return TYPE_NUM
	} else {
		return TYPE_PART
	}
}

var dirs = [8][2]int16{
	{-1, 1},
	{0, 1},
	{1, 1},
	{-1, 0},
	{1, 0},
	{-1, -1},
	{0, -1},
	{1, -1}}

type Part struct {
	PartType byte
	X, Y     int16
}

type Pair struct {
	Part    Part
	Matches []Match
}

func NewPart(t byte, x, y int16) *Part {
	return &Part{t, x, y}
}

func FindObjects(field *aoc_utils.ByteField) ([]Part, []Match) {
	parts := make([]Part, 0)
	matches := make([]Match, 0)
	var currentMatch Match = Match{}
	currentMatch.Value = -1
	endNum := func() {
		if currentMatch.Value < 0 {
			return
		}
		matches = append(matches, currentMatch)
		currentMatch = Match{}
		currentMatch.Value = -1
	}
	var cellType uint8
	for y := range int16((*field).Height) {
		for x := range int16((*field).Width) {
			cellType = GetType((*field).Field[y][x])
			if cellType == TYPE_NUM {
				if currentMatch.Value < 0 {
					currentMatch.Value = int((*field).Field[y][x] - '0')
					currentMatch.StartX = x
					currentMatch.Y = y
				} else {
					currentMatch.Value = (currentMatch.Value * 10) + int((*field).Field[y][x]-'0')
					currentMatch.EndX = x
				}
				continue
			}
			endNum()
			if cellType == TYPE_PART {
				part := Part{}
				part.PartType = (*field).Field[y][x]
				part.X = x
				part.Y = y
				parts = append(parts, part)
			}
		}
		// catch case that the last value is a number and didn't close
		endNum()
	}
	return parts, matches
}

func PairObjects(parts []Part, matches []Match) []Pair {
	data := make([]Pair, 0)
	var testX, testY int16
	for partIdx := range parts {
		group := make([]Match, 0)
		for _, dir := range dirs {
			testX = parts[partIdx].X + dir[1]
			testY = parts[partIdx].Y + dir[0]
			for _, match := range matches {
				//fmt.Printf("%d %d %d %d %d\n", testX, testY, match.Y, match.StartX, match.EndX)
				if match.Y == testY && match.StartX <= testX && match.EndX >= testX && !slices.Contains(group, match) {
					group = append(group, match)
				}
			}
		}
		if len(group) > 0 {
			data = append(data, Pair{parts[partIdx], group})
		}
	}
	return data
}

func Part1(in io.Reader) int {
	field, _ := aoc_utils.LoadField(in)
	parts, matches := FindObjects(field)
	pairs := PairObjects(parts, matches)
	summe := 0
	for pairIdx := range pairs {
		//fmt.Printf("Part %v has %v matches\n", (pairs[pairIdx].Part), pairs[pairIdx].Matches)
		for _, match := range pairs[pairIdx].Matches {
			summe += match.Value
		}
	}
	return summe
}
