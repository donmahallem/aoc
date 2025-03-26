package day12_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day12"
	"github.com/donmahallem/aoc/utils"
)

func TestSortHorizontal(t *testing.T) {
	test := [][2]int{{3, 3}, {3, 1}, {1, 1}}
	expected := [][2]int{{1, 1}, {3, 1}, {3, 3}}
	slices.SortFunc(test, day12.SortHorizontal)
	if !slices.Equal(test, expected) {
		t.Errorf(`Expected %v to match %v`, test, expected)
	}
}
func TestSortVertical(t *testing.T) {
	test := [][2]int{{3, 3}, {3, 1}, {1, 1}, {4, 1}}
	expected := [][2]int{{1, 1}, {3, 1}, {4, 1}, {3, 3}}
	slices.SortFunc(test, day12.SortVertical)
	if !slices.Equal(test, expected) {
		t.Errorf(`Expected %v to match %v`, test, expected)
	}
}

func TestCountStraightEdges(t *testing.T) {
	test, _ := utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 0, 4)
	result := day12.CountStraightEdges(neighbours)
	if result != 8 {
		t.Errorf(`Expected %d to match %d`, result, 8)
	}
}

func TestCountStraightEdgesHorizontal(t *testing.T) {
	test, _ := utils.LoadField(strings.NewReader(testData))
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
	test, _ := utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 0, 4)
	result := day12.CountStraightEdgesVertical(neighbours)
	if result != 2 {
		t.Errorf(`Expected %d to match %d`, result, 2)
	}
}

func TestCountStraightEdgesVertical_R(t *testing.T) {
	test, _ := utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 0, 0)
	result := day12.CountStraightEdgesVertical(neighbours)
	if result != 5 {
		t.Errorf(`Expected %d to match %d`, result, 5)
	}
}
func TestCountStraightEdgesHorizontal_R(t *testing.T) {
	test, _ := utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 0, 0)
	result := day12.CountStraightEdgesHorizontal(neighbours)
	if result != 5 {
		t.Errorf(`Expected %d to match %d`, result, 5)
	}
}

func TestCountStraightEdgesVertical_7_0(t *testing.T) {
	test, _ := utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 7, 0)
	result := day12.CountStraightEdgesVertical(neighbours)
	if result != 3 {
		t.Errorf(`Expected %d to match %d`, result, 3)
	}
}

func TestCountStraightEdgesHorizontal_7_0(t *testing.T) {
	test, _ := utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 7, 0)
	result := day12.CountStraightEdgesHorizontal(neighbours)
	if result != 3 {
		t.Errorf(`Expected %d to match %d`, result, 3)
	}
}

func TestCountStraightEdgesHorizontal_V(t *testing.T) {
	test, _ := utils.LoadField(strings.NewReader(testData))
	neighbours := day12.FindNeighbours(test, 2, 0)
	result := day12.CountStraightEdgesHorizontal(neighbours)
	if result != 5 {
		t.Errorf(`Expected %d to match %d`, result, 5)
	}
}

func TestCountStraightEdgesVertical_ReverseC(t *testing.T) {
	test := [][2]int{{0, 0}, {0, 1}, {1, 1}, {2, 0}, {2, 1}}
	result := day12.CountStraightEdgesVertical(test)
	if result != 4 {
		t.Errorf(`Expected %d to match %d`, result, 4)
	}
}

func TestCountStraightEdgesVertical_ReverseCBuckle(t *testing.T) {
	test := [][2]int{{0, 0}, {0, 1}, {1, 1}, {2, 0}, {2, 1}, {1, 2}}
	result := day12.CountStraightEdgesVertical(test)
	if result != 6 {
		t.Errorf(`Expected %d to match %d`, result, 6)
	}
}
