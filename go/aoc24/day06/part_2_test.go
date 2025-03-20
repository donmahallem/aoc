package day06

import (
	"strings"
	"testing"
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
	testField, guard, _ := readSource(strings.NewReader(testData))
	testField.field[8][1] = true
	if testLoop(&testField, guard) {
		t.Errorf(`Expected %s to match`, "a")
	}
}
