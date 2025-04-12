package day02_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day02"
)

const testData string = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPart1(t *testing.T) {
	result := day02.Part1(strings.NewReader(testData))
	if result != 2 {
		t.Errorf(`Expected %d to be %d`, result, 2)
	}
}

func BenchmarkPart1(b *testing.B) {
	for b.Loop() {
		day02.Part1(strings.NewReader(testData))
	}
}
