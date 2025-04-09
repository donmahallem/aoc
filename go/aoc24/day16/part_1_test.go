package day16_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day16"
)

const testData1 string = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`
const testData2 string = `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`

func TestParseInput(t *testing.T) {
	field, start, end := day16.ParseInput(strings.NewReader(testData1))
	if len(field) != 15 {
		t.Errorf(`Expected %d to match %d`, len(field), 15)
	}
	if start.Y != 13 || start.X != 1 {
		t.Errorf(`Expected %v to match [4,4]`, start)
	}
	if end.Y != 1 || end.X != 13 {
		t.Errorf(`Expected %v to match [4,4]`, end)
	}
}

func fieldToCsv(f *day16.Field) {
	p := day16.Point{}
	for y := range len(*f) {
		p.Y = int16(y)
		for x := range len((*f)[y]) {
			p.X = int16(x)
			if x > 0 {
				fmt.Print(",")
			}
			fmt.Printf("%d", (*f)[y][x])
		}
		fmt.Println()
	}
}

func TestFindShortestPath(t *testing.T) {
	field, start, end := day16.ParseInput(strings.NewReader(testData1))
	day16.CalculatePathValues(&field, &start)
	if field[end.Y][end.X] != 7036 {
		t.Errorf(`Expected %d to match 7036`, field[end.Y][end.X])
	}
}

func TestPart1_testData1(t *testing.T) {
	result := day16.Part1(strings.NewReader(testData1))
	if result != 7036 {
		t.Errorf(`Expected %d to match 7036`, result)
	}
}

func TestPart1_testData2(t *testing.T) {
	result := day16.Part1(strings.NewReader(testData2))
	if result != 11048 {
		t.Errorf(`Expected %d to match 11048`, result)
	}
}
