package day03

import (
	"slices"

	"github.com/donmahallem/aoc/go/aoc_utils"
	"github.com/donmahallem/aoc/go/aoc_utils/bytes"
)

const (
	typeNum uint8 = iota
	typePart
	typeSpace
)

// Number value and Position inside the grid
type number struct {
	Value   int // Value of the Number
	StartX, // X Start Position of the Number
	EndX, // X End Position (inclusive) of the Number
	Y int16 // Y Position of the Number
}

func newNumber(t int, startX, endX, y int16) *number {
	return &number{t, startX, endX, y}
}

func getType(b byte) uint8 {
	if b == '.' {
		return typeSpace
	} else if bytes.ByteIsNumber(b) {
		return typeNum
	} else {
		return typePart
	}
}

type field = aoc_utils.ByteField[int16, byte]

var dirs = [8][2]int16{
	{-1, 1},
	{0, 1},
	{1, 1},
	{-1, 0},
	{1, 0},
	{-1, -1},
	{0, -1},
	{1, -1}}

type part struct {
	// Part Character
	PartType byte
	// X Position of the part in the grid
	X,
	// Y Position of the part in the grid
	Y int16
}

// Pair of Part to neighbouring numbers
type pair struct {
	// Part info
	Part part
	// Numbers touching the part
	Matches []number
}

func newPart(t byte, x, y int16) *part {
	return &part{t, x, y}
}

func sortNumbers(a number, b number) int {
	if a.Y == b.Y {
		return int(a.StartX - b.EndX)
	}
	return int(a.Y - b.Y)
}

// Searches the Field for Parts and Numbers
func findObjects(f field) ([]part, []number) {
	parts := make([]part, 0, 16)
	matches := make([]number, 0, 16)
	var currentMatch number = number{}
	currentMatch.Value = -1
	// Checks if a number was encountered.
	// If so it finishes up the current and
	// declares a new one
	endNum := func() {
		if currentMatch.Value < 0 {
			return
		}
		matches = append(matches, currentMatch)
		currentMatch = number{}
		currentMatch.Value = -1
	}
	var cellType uint8
	for y := int16(0); y < f.Height; y++ {
		for x := int16(0); x < f.Width; x++ {
			cellValue := f.Get(x, y)
			cellType = getType(cellValue)
			if cellType == typeNum {
				if currentMatch.Value < 0 {
					currentMatch.Value = int(cellValue - '0')
					currentMatch.StartX = x
					currentMatch.EndX = x
					currentMatch.Y = y
				} else {
					currentMatch.Value = (currentMatch.Value * 10) + int(cellValue-'0')
					currentMatch.EndX = x
				}
				continue
			}
			endNum()
			if cellType == typePart {
				p := part{}
				p.PartType = cellValue
				p.X = x
				p.Y = y
				parts = append(parts, p)
			}
		}
		// catch case that the last value is a number and didn't close
		endNum()
	}
	return parts, matches
}

// Search for
func pairObjects(parts []part, matches []number) []pair {
	data := make([]pair, 0, len(parts))
	var testX, testY int16
	for partIdx := range len(parts) {
		group := make([]number, 0)
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
			p := pair{}
			p.Part = parts[partIdx]
			p.Matches = group
			data = append(data, p)
		}
	}
	return data
}
