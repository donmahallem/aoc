package day13_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day13"
)

func TestPart2(t *testing.T) {
	if result := day13.Part2(strings.NewReader(testData)); result != 875318608908 {
		t.Errorf(`Expected %d to contain %d`, result, 875318608908)
	}
}
