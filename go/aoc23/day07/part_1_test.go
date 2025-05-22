package day07_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day07"
)

const testData string = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestOutOfBoundsShouldBeInside(t *testing.T) {
	var expected uint = 6440
	reader := strings.NewReader(testData)
	if res := day07.Part1(reader); res != expected {
		t.Errorf(`Expected %v to match %v`, res, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day07.Part1(reader)
	}
}
