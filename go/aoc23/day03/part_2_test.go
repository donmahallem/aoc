package day03_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day03"
)

func TestPart2(t *testing.T) {
	result := day03.Part2(strings.NewReader(testData))
	if result != 467835 {
		t.Errorf(`Expected %d to be %d`, result, 467835)
	}
}

func BenchmarkPart2(b *testing.B) {
	for b.Loop() {
		day03.Part2(strings.NewReader(testData))
	}
}
