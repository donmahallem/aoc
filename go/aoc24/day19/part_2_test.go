package day19_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day19"
)

func TestPart2(t *testing.T) {
	test, err := day19.Part2(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if test != 16 {
		t.Errorf(`Expected %d to match 16`, test)
	}
}

func BenchmarkPart2(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day19.Part2(data)
	}
}
