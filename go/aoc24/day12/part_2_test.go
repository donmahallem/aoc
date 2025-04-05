package day12_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day12"
	"github.com/donmahallem/aoc/aoc_utils"
)

func TestSortHorizontal(t *testing.T) {
	test := []day12.Point{{Y: 3, X: 3}, {Y: 3, X: 1}, {Y: 1, X: 1}}
	expected := []day12.Point{{Y: 1, X: 1}, {Y: 3, X: 1}, {Y: 3, X: 3}}
	slices.SortFunc(test, day12.SortHorizontal)
	if !slices.Equal(test, expected) {
		t.Errorf(`Expected %v to match %v`, test, expected)
	}
}
func TestSortVertical(t *testing.T) {
	test := []day12.Point{{Y: 3, X: 3}, {Y: 3, X: 1}, {Y: 1, X: 1}, {Y: 4, X: 1}}
	expected := []day12.Point{{Y: 1, X: 1}, {Y: 3, X: 1}, {Y: 4, X: 1}, {Y: 3, X: 3}}
	slices.SortFunc(test, day12.SortVertical)
	if !slices.Equal(test, expected) {
		t.Errorf(`Expected %v to match %v`, test, expected)
	}
}

func TestCountStraightEdges(t *testing.T) {
	test, _ := aoc_utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 0, 4)
	result := day12.CountStraightEdges(neighbours)
	if result != 4 {
		t.Errorf(`Expected %d to match %d`, result, 4)
	}
}

func TestCountStraightEdgesHorizontal(t *testing.T) {
	test, _ := aoc_utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 0, 4)
	result := day12.CountStraightEdgesHorizontal(neighbours)
	if result != 2 {
		t.Errorf(`Expected %d to match %d`, result, 2)
	}
	neighbours = day12.FindNeighbours(test, 0, 0)
	result = day12.CountStraightEdgesHorizontal(neighbours)
	if result != 5 {
		t.Errorf(`Expected %d to match %d`, result, 5)
	}
}

func TestCountStraightEdgesVertical_0_4(t *testing.T) {
	test, _ := aoc_utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 0, 4)
	result := day12.CountStraightEdgesVertical(neighbours)
	if result != 2 {
		t.Errorf(`Expected %d to match %d`, result, 2)
	}
}

func TestCountStraightEdgesVertical_R(t *testing.T) {
	test, _ := aoc_utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 0, 0)
	result := day12.CountStraightEdgesVertical(neighbours)
	if result != 5 {
		t.Errorf(`Expected %d to match %d`, result, 5)
	}
}
func TestCountStraightEdgesHorizontal_R(t *testing.T) {
	test, _ := aoc_utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 0, 0)
	result := day12.CountStraightEdgesHorizontal(neighbours)
	if result != 5 {
		t.Errorf(`Expected %d to match %d`, result, 5)
	}
}

func TestCountStraightEdgesVertical_7_0(t *testing.T) {
	test, _ := aoc_utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 7, 0)
	result := day12.CountStraightEdgesVertical(neighbours)
	if result != 3 {
		t.Errorf(`Expected %d to match %d`, result, 3)
	}
}

func TestCountStraightEdgesHorizontal_7_0(t *testing.T) {
	test, _ := aoc_utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 7, 0)
	result := day12.CountStraightEdgesHorizontal(neighbours)
	if result != 3 {
		t.Errorf(`Expected %d to match %d`, result, 3)
	}
}

func TestCountStraightEdgesHorizontal_V(t *testing.T) {
	test, _ := aoc_utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 2, 0)
	result := day12.CountStraightEdgesHorizontal(neighbours)
	if result != 5 {
		t.Errorf(`Expected %d to match %d`, result, 5)
	}
}

func TestCountStraightEdgesVertical_ReverseC(t *testing.T) {
	test := []day12.Point{{Y: 0, X: 0}, {Y: 0, X: 1}, {Y: 1, X: 1}, {Y: 2, X: 0}, {Y: 2, X: 1}}
	result := day12.CountStraightEdgesVertical(test)
	if result != 4 {
		t.Errorf(`Expected %d to match %d`, result, 4)
	}
}

func TestCountStraightEdgesVertical_ReverseCBuckle(t *testing.T) {
	test := []day12.Point{{Y: 0, X: 0}, {Y: 0, X: 1}, {Y: 1, X: 1}, {Y: 2, X: 0}, {Y: 2, X: 1}, {Y: 1, X: 2}}
	result := day12.CountStraightEdgesVertical(test)
	if result != 6 {
		t.Errorf(`Expected %d to match %d`, result, 6)
	}
}

func TestPart2(t *testing.T) {
	if result := day12.Part2(strings.NewReader(testData)); result != 1206 {
		t.Errorf(`Expected %d to contain %d`, result, 1206)
	}
}
