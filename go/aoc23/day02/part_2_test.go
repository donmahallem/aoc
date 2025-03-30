package day02_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day02"
)

const testData2 string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestCalculateMinBlock(t *testing.T) {
	data := []byte(testData[0])
	_, blocks := day02.ParseLine(&data)
	if res := day02.CalculateMinBlock(&blocks); res != 48 {
		t.Errorf(`Expected %d to be %d`, res, 48)
	}
}

func TestPart2(t *testing.T) {
	result := day02.Part2(strings.NewReader(testData2))
	if result != 2286 {
		t.Errorf(`Expected %d to be %d`, result, 2286)
	}
}
