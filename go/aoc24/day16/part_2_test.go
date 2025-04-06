package day16_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day16"
)

func TestPart2(t *testing.T) {
	result := day16.Part2(strings.NewReader(testData1))
	if result != 45 {
		t.Errorf(`Expected %d to match 45`, result)
	}
}
