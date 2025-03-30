package day04_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day04"
)

func TestPart2(t *testing.T) {
	result := day04.Part2(strings.NewReader(testData))
	if result != 30 {
		t.Errorf(`Expected winners to have a length of %d. Not %d`, 30, result)
	}
}
