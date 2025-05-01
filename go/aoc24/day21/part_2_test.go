package day21_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day21"
)

func TestPart2(t *testing.T) {
	test := day21.Part2(strings.NewReader(testData))
	if test != 724554 {
		t.Errorf(`Expected %d to match 724554`, test)
	}
}
