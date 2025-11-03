package day01_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day01"
)

func Test24Day01Part2(t *testing.T) {

	if result := day01.Part2(strings.NewReader(testData)); result != 31 {
		t.Errorf(`Expected %d to match %d`, result, 31)
	}
}

func BenchmarkPart2(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day01.Part2(reader)
	}
}
