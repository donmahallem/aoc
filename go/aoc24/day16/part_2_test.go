package day16_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day16"
)

func TestPart2_1(t *testing.T) {
	result := day16.Part2(strings.NewReader(testData1))
	if result != 45 {
		t.Errorf(`Expected %d to match 45`, result)
	}
}
func TestPart2_2(t *testing.T) {
	result := day16.Part2(strings.NewReader(testData2))
	if result != 64 {
		t.Errorf(`Expected %d to match 64`, result)
	}
}
