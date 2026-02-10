package day12

import (
	"io"
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func TestSortHorizontal(t *testing.T) {
	test := []point{{Y: 3, X: 3}, {Y: 3, X: 1}, {Y: 1, X: 1}}
	expected := []point{{Y: 1, X: 1}, {Y: 3, X: 1}, {Y: 3, X: 3}}
	slices.SortFunc(test, sortHorizontal)
	if !slices.Equal(test, expected) {
		t.Errorf(`Expected %v to match %v`, test, expected)
	}
}
func TestSortVertical(t *testing.T) {
	test := []point{{Y: 3, X: 3}, {Y: 3, X: 1}, {Y: 1, X: 1}, {Y: 4, X: 1}}
	expected := []point{{Y: 1, X: 1}, {Y: 3, X: 1}, {Y: 4, X: 1}, {Y: 3, X: 3}}
	slices.SortFunc(test, sortVertical)
	if !slices.Equal(test, expected) {
		t.Errorf(`Expected %v to match %v`, test, expected)
	}
}

func TestCountStraightEdges(t *testing.T) {
	test, _ := aoc_utils.LoadField[int16, byte](strings.NewReader(testData))
	neighbours := findNeighbours(*test, 0, 4)
	result := countStraightEdges(neighbours)
	if result != 4 {
		t.Errorf(`Expected %d to match %d`, result, 4)
	}
}

func TestCountStraightEdgesHorizontal(t *testing.T) {
	test, _ := aoc_utils.LoadField[int16, byte](strings.NewReader(testData))
	neighbours := findNeighbours(*test, 0, 4)
	result := countStraightEdgesHorizontal(neighbours)
	if result != 2 {
		t.Errorf(`Expected %d to match %d`, result, 2)
	}
	neighbours = findNeighbours(*test, 0, 0)
	result = countStraightEdgesHorizontal(neighbours)
	if result != 5 {
		t.Errorf(`Expected %d to match %d`, result, 5)
	}
}

func TestCountStraightEdgesVertical_0_4(t *testing.T) {
	test, _ := aoc_utils.LoadField[int16, byte](strings.NewReader(testData))
	neighbours := findNeighbours(*test, 0, 4)
	result := countStraightEdgesVertical(neighbours)
	if result != 2 {
		t.Errorf(`Expected %d to match %d`, result, 2)
	}
}

func TestCountStraightEdgesVertical_R(t *testing.T) {
	test, _ := aoc_utils.LoadField[int16, byte](strings.NewReader(testData))
	neighbours := findNeighbours(*test, 0, 0)
	result := countStraightEdgesVertical(neighbours)
	if result != 5 {
		t.Errorf(`Expected %d to match %d`, result, 5)
	}
}
func TestCountStraightEdgesHorizontal_R(t *testing.T) {
	test, _ := aoc_utils.LoadField[int16, byte](strings.NewReader(testData))
	neighbours := findNeighbours(*test, 0, 0)
	result := countStraightEdgesHorizontal(neighbours)
	if result != 5 {
		t.Errorf(`Expected %d to match %d`, result, 5)
	}
}

func TestCountStraightEdgesVertical_7_0(t *testing.T) {
	test, _ := aoc_utils.LoadField[int16, byte](strings.NewReader(testData))
	neighbours := findNeighbours(*test, 7, 0)
	result := countStraightEdgesVertical(neighbours)
	if result != 3 {
		t.Errorf(`Expected %d to match %d`, result, 3)
	}
}

func TestCountStraightEdgesHorizontal_7_0(t *testing.T) {
	test, _ := aoc_utils.LoadField[int16, byte](strings.NewReader(testData))
	neighbours := findNeighbours(*test, 7, 0)
	result := countStraightEdgesHorizontal(neighbours)
	if result != 3 {
		t.Errorf(`Expected %d to match %d`, result, 3)
	}
}

func TestCountStraightEdgesHorizontal_V(t *testing.T) {
	test, _ := aoc_utils.LoadField[int16, byte](strings.NewReader(testData))
	neighbours := findNeighbours(*test, 2, 0)
	result := countStraightEdgesHorizontal(neighbours)
	if result != 5 {
		t.Errorf(`Expected %d to match %d`, result, 5)
	}
}

func TestCountStraightEdgesVertical_ReverseC(t *testing.T) {
	test := []point{{Y: 0, X: 0}, {Y: 0, X: 1}, {Y: 1, X: 1}, {Y: 2, X: 0}, {Y: 2, X: 1}}
	result := countStraightEdgesVertical(test)
	if result != 4 {
		t.Errorf(`Expected %d to match %d`, result, 4)
	}
}

func TestCountStraightEdgesVertical_ReverseCBuckle(t *testing.T) {
	test := []point{{Y: 0, X: 0}, {Y: 0, X: 1}, {Y: 1, X: 1}, {Y: 2, X: 0}, {Y: 2, X: 1}, {Y: 1, X: 2}}
	result := countStraightEdgesVertical(test)
	if result != 6 {
		t.Errorf(`Expected %d to match %d`, result, 6)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 1206 {
		t.Errorf(`Expected %d to contain %d`, result, 1206)
	}
}

func BenchmarkPart2(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		Part2(data)
	}
}
