package day12_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day12"
	"github.com/donmahallem/aoc/aoc_utils"
)

const testData string = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func TestCountEdges(t *testing.T) {
	test := [][2]int{{2, 2}, {2, 3}, {3, 3}}
	result := day12.CountEdges(test)
	if result != 8 {
		t.Errorf(`Expected %d to match %d`, result, 8)
	}
}

func TestFindNeighbours(t *testing.T) {
	test, _ := aoc_utils.LoadField(strings.NewReader(testData))
	result := day12.FindNeighbours(test, 0, 4)
	expected := [][2]int{{0, 5}, {0, 4}, {1, 4}, {1, 5}}
	if len(result) != 4 {
		t.Errorf(`Expected result to have a length of 4 not %d`, len(result))
	}
	for _, obj := range expected {

		if !slices.Contains(result, obj) {
			t.Errorf(`Expected %v to contain %v`, result, obj)
		}
	}
}

func TestPart1(t *testing.T) {
	if result := day12.Part1(strings.NewReader(testData)); result != 1930 {
		t.Errorf(`Expected %d to contain %d`, result, 1930)
	}
}
