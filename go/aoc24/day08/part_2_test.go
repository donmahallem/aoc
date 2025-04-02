package day08_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day08"
)

func Test24Day08Part2(t *testing.T) {

	if result := day08.Part2(strings.NewReader(testData)); result != 34 {
		t.Errorf(`Expected %d to match %d`, result, 34)
	}
}
