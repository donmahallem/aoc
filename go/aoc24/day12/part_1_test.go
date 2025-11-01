package day12_test

import (
	"io"
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
	test := []day12.Point{{Y: 2, X: 2}, {Y: 2, X: 3}, {Y: 3, X: 3}}
	result := day12.CountEdges(test)
	if result != 8 {
		t.Errorf(`Expected %d to match %d`, result, 8)
	}
}

func TestFindNeighbours(t *testing.T) {
	test, _ := aoc_utils.LoadField[int16, byte](strings.NewReader(testData))
	result := day12.FindNeighbours(*test, 0, 4)
	expected := []day12.Point{{Y: 0, X: 5}, {Y: 0, X: 4}, {Y: 1, X: 4}, {Y: 1, X: 5}}
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

func BenchmarkPart1(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day12.Part1(data)
	}
}
