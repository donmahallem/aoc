package day01_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day01"
)

const testDataPart2 string = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestParseFilePart2(t *testing.T) {
	expected := 281
	reader := strings.NewReader(testDataPart2)
	if res := day01.Part2(reader); res != expected {
		t.Errorf(`Expected %v to match %v`, res, expected)
	}
}

func BenchmarkPart2(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day01.Part2(reader)
	}
}
