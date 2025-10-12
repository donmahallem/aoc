package day10_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day10"
)

var testData string = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

func TestPart1(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 8
		reader := strings.NewReader(testData)
		if res := day10.Part1(reader); res != expected {
			t.Errorf(`Expected %d to be %d`, res, expected)
		}
	})
}
