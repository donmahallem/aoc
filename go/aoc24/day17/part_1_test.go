package day17_test

import (
	_ "embed"
	"io"
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day17"
)

//go:embed sample1.txt
var testData1 string

func TestPart1_testData1(t *testing.T) {
	result, err := day17.Part1(strings.NewReader(testData1))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []int{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}
	if !slices.Equal(result, expected) {
		t.Errorf(`Expected %v to match %v`, result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	data := strings.NewReader(testData2)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day17.Part1(data)
	}
}
