package day01_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day01"
)

const testData string = `3   4
4   3
2   5
1   3
3   9
3   3`

func Test24Day01Part1(t *testing.T) {

	if result := day01.Part1(strings.NewReader(testData)); result != 11 {
		t.Errorf(`Expected %d to match %d`, result, 11)
	}
}

func BenchmarkPart1(b *testing.B) {
	for b.Loop() {
		day01.Part1(strings.NewReader(testData))
	}
}
