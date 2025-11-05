package day06_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day06"
)

var testData string = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestPart2(t *testing.T) {
	if res := day06.Part2(strings.NewReader(testData)); res != 6 {
		t.Errorf(`Expected %d to match %d`, res, 6)
	}
}

func BenchmarkPart2(b *testing.B) {
	testData := strings.NewReader(testData)
	for b.Loop() {
		testData.Seek(0, io.SeekStart)
		day06.Part2(testData)
	}
}
