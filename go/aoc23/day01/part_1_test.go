package day01_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day01"
)

const testData string = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func TestOutOfBoundsShouldBeInside(t *testing.T) {
	expected := 142
	reader := strings.NewReader(testData)
	if res := day01.Part1(reader); res != expected {
		t.Errorf(`Expected %v to match %v`, res, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day01.Part1(reader)
	}
}
