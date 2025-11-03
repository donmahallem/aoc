package day13_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day13"
)

func TestPart2(t *testing.T) {
	if result := day13.Part2(strings.NewReader(testData)); result != 875318608908 {
		t.Errorf(`Expected %d to contain %d`, result, 875318608908)
	}
}

func BenchmarkPart2(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day13.Part2(data)
	}
}
