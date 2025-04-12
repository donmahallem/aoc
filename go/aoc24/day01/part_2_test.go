package day01_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day01"
)

func Test24Day01Part2(t *testing.T) {

	if result := day01.Part2(strings.NewReader(testData)); result != 31 {
		t.Errorf(`Expected %d to match %d`, result, 31)
	}
}

func BenchmarkPart2(b *testing.B) {
	for b.Loop() {
		day01.Part2(strings.NewReader(testData))
	}
}
