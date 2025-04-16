package day20_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day20"
)

const testData string = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

var COUNT_CHEATS_TESTS [][2]int = [][2]int{{2, 14},
	{4, 14},
	{6, 2},
	{8, 4},
	{10, 2},
	{12, 3},
	{20, 1},
	{36, 1},
	{38, 1},
	{40, 1},
	{64, 1}}

func TestParseInput(t *testing.T) {
	raceWay := day20.ParseInput(strings.NewReader(testData))
	if len(*raceWay) != 85 {
		t.Errorf(`Expected %d to match 85`, len(*raceWay))
	}
}
func TestCountCheats(t *testing.T) {
	patterns := day20.ParseInput(strings.NewReader(testData))
	if result := day20.CountCheats(patterns, 2); result != 44 {
		t.Errorf(`Expected %d to save %d steps. Not %d`, 2, 44, result)
	}
}

func BenchmarkParseInput(b *testing.B) {
	for b.Loop() {
		day20.ParseInput(strings.NewReader(testData))
	}
}

func BenchmarkCountCheats(b *testing.B) {
	patterns := day20.ParseInput(strings.NewReader(testData))
	for b.Loop() {
		day20.CountCheats(patterns, 2)
	}
}
