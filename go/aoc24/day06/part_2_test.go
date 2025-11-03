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

func TestLoop(t *testing.T) {
	testField, guard, _ := day06.ReadSource(strings.NewReader(testData))
	testField.Field[8][1] = true
	if !day06.TestLoop(&testField, guard) {
		t.Errorf(`Expected to be valid`)
	}
}

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
