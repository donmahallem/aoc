package day01_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day01"
)

const testData string = `3   4
4   3
2   5
1   3
3   9
3   3`

func Test24Day01Part1(t *testing.T) {

	result, err := day01.Part1(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 11 {
		t.Errorf(`Expected %d to match %d`, result, 11)
	}
}

func BenchmarkPart1(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day01.Part1(reader)
	}
}
