package day17_test

import (
	_ "embed"
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day17"
)

//go:embed sample2.txt
var testData2 string

func TestPart2_testData2(t *testing.T) {
	result, err := day17.Part2(strings.NewReader(testData2))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 117440 {
		t.Errorf(`Expected %d to match 117440`, result)
	}
}

func BenchmarkPart2(b *testing.B) {
	data := strings.NewReader(testData2)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day17.Part2(data)
	}
}
