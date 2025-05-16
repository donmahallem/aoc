package day09_test

import (
	"io"
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day09"
)

const testData string = "2333133121414131402"

func TestOutOfBoundsShouldBeInside(t *testing.T) {
	test := []byte{'1', '2', '3', '4', '5'}
	expected := []int16{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}
	if res := day09.ConvertInput(&test); !slices.Equal(res, expected) {
		t.Errorf(`Expected %v to match %v`, res, expected)
	}
}
func TestOutOfBoundsShouldBeInside2(t *testing.T) {
	test := []int16{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}
	expected := []int16{0, 2, 2, 1, 1, 1, 2, 2, 2, -1, -1, -1, -1, -1, -1}
	day09.CompactData(&test)
	if !slices.Equal(test, expected) {
		t.Errorf(`Expected %v to match %v`, test, expected)
	}
}

func TestPart1(t *testing.T) {
	data := strings.NewReader(testData)
	result := day09.Part1(data)
	if result != 1928 {
		t.Errorf("Expected result to be 1928. Got %d", result)
	}
}

func BenchmarkPart1(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day09.Part1(data)
	}
}
