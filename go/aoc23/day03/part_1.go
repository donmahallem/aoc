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

// Number value and Position inside the grid
type Number struct {
	Value   int // Value of the Number
	StartX, // X Start Position of the Number
	EndX, // X End Position (inclusive) of the Number
	Y int16 // Y Position of the Number
}

func NewNumber(t int, startX, endX, y int16) *Number {
	return &Number{t, startX, endX, y}
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
	// Part Character
	PartType byte
	// X Position of the part in the grid
	X,
	// Y Position of the part in the grid
	Y int16
}

// Pair of Part to neighbouring numbers
type Pair struct {
	// Part info
	Part Part
	// Numbers touching the part
	Matches []Number
}

func NewPart(t byte, x, y int16) *Part {
	return &Part{t, x, y}
}

func SortNumbers(a Number, b Number) int {
	if a.Y == b.Y {
		return int(a.StartX - b.EndX)
	}
	return int(a.Y - b.Y)
}

// Searches the Field for Parts and Numbers
func FindObjects(field *aoc_utils.ByteField) ([]Part, []Number) {
	parts := make([]Part, 0)
	matches := make([]Number, 0)
	var currentMatch Number = Number{}
	currentMatch.Value = -1
	// Checks if a number was encountered.
	// If so it finishes up the current and
	// declares a new one
	endNum := func() {
		if currentMatch.Value < 0 {
			return
		}
		matches = append(matches, currentMatch)
		currentMatch = Number{}
		currentMatch.Value = -1
	}
	var cellType uint8
	for y := range int16(field.Height) {
		for x := range int16(field.Width) {
			cellType = GetType(field.Field[y][x])
			if cellType == TYPE_NUM {
				if currentMatch.Value < 0 {
					currentMatch.Value = int(field.Field[y][x] - '0')
					currentMatch.StartX = x
					currentMatch.EndX = x
					currentMatch.Y = y
				} else {
					currentMatch.Value = (currentMatch.Value * 10) + int(field.Field[y][x]-'0')
					currentMatch.EndX = x
				}
				continue
			}
			endNum()
			if cellType == TYPE_PART {
				part := Part{}
				part.PartType = field.Field[y][x]
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

// Search for
func PairObjects(parts []Part, matches []Number) []Pair {
	data := make([]Pair, 0)
	var testX, testY int16
	for partIdx := range len(parts) {
		group := make([]Number, 0)
		for dirIdx := range len(dirs) {
			testX = parts[partIdx].X + dirs[dirIdx][1]
			testY = parts[partIdx].Y + dirs[dirIdx][0]
			for _, match := range matches {
				if match.Y == testY && match.StartX <= testX && match.EndX >= testX && !slices.Contains(group, match) {
					group = append(group, match)
				}
			}
		}
		if len(group) > 0 {
			pair := Pair{}
			pair.Part = parts[partIdx]
			pair.Matches = group
			data = append(data, pair)
		}
	}
	return data
}

func Part1(in io.Reader) int {
	field, _ := aoc_utils.LoadField(in)
	parts, matches := FindObjects(field)
	pairs := PairObjects(parts, matches)
	summe := 0
	for pairIdx := range len(pairs) {
		for matchIdx := range len(pairs[pairIdx].Matches) {
			summe += pairs[pairIdx].Matches[matchIdx].Value
		}
	}
	return summe
}
