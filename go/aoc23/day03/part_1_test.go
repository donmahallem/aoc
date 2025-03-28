package day03_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day03"
	"github.com/donmahallem/aoc/aoc_utils"
)

var testData = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestCheckForPair(t *testing.T) {
	result, _ := aoc_utils.LoadField(strings.NewReader(testData))
	testMatch := day03.Match{}
	testMatch.Start.X = 0
	testMatch.Start.Y = 0
	testMatch.End.X = 2
	testMatch.End.Y = 0

	if !day03.CheckForPair(result, &testMatch) {
		t.Errorf(`Expected to be true`)
	}
}

func TestParseField(t *testing.T) {
	result := day03.ParseField(strings.NewReader(testData))
	if len(result) != 10 {
		t.Errorf(`Expected %d to be %d`, len(result), 10)
	}
}

func TestPart1(t *testing.T) {
	result := day03.Part1(strings.NewReader(testData))
	if result != 4361 {
		t.Errorf(`Expected %d to be %d`, result, 4361)
	}
}
