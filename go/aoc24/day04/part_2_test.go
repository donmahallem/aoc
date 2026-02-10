package day04_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day04"
)

func Test24Day04Part2(t *testing.T) {

	result, err := day04.Part2(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 9 {
		t.Errorf(`Expected %d to match %d`, result, 9)
	}
}

func BenchmarkPart2(b *testing.B) {
	testData := strings.NewReader(testData)
	for b.Loop() {
		testData.Seek(0, io.SeekStart)
		day04.Part2(testData)
	}
}
