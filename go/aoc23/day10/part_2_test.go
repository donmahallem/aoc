package day10_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day10"
)

var testData2 string = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

func TestPart2(t *testing.T) {
	t.Run("testData2", func(t *testing.T) {
		const expected int = 4
		reader := strings.NewReader(testData2)
		if res := day10.Part2(reader); res != expected {
			t.Errorf(`Expected %d to be %d`, res, expected)
		}
	})
}
