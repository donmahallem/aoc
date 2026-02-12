package day12

import (
	"slices"
	"strings"
	"testing"

	_ "embed"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

//go:embed sample1.txt
var testData string

func TestCountEdges(t *testing.T) {
	test := []point{{Y: 2, X: 2}, {Y: 2, X: 3}, {Y: 3, X: 3}}
	result := countEdges(test)
	if result != 8 {
		t.Errorf(`Expected %d to match %d`, result, 8)
	}
}

func TestFindNeighbours(t *testing.T) {
	test, _ := aoc_utils.LoadField[int16, byte](strings.NewReader(testData))
	result := findNeighbours(*test, 0, 4)
	expected := []point{{Y: 0, X: 5}, {Y: 0, X: 4}, {Y: 1, X: 4}, {Y: 1, X: 5}}
	if len(result) != 4 {
		t.Errorf(`Expected result to have a length of 4 not %d`, len(result))
	}
	for _, obj := range expected {

		if !slices.Contains(result, obj) {
			t.Errorf(`Expected %v to contain %v`, result, obj)
		}
	}
}
