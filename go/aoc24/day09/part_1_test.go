package day09

import (
	"io"
	"slices"
	"strings"
	"testing"
)

const testData string = "2333133121414131402"

func TestOutOfBoundsShouldBeInside(t *testing.T) {
	test := []byte{'1', '2', '3', '4', '5'}
	expected := []int16{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}
	if res, err := convertInput(test); err != nil || !slices.Equal(res, expected) {
		t.Errorf(`Expected %v to match %v`, res, expected)
	}
}
func TestOutOfBoundsShouldBeInside2(t *testing.T) {
	test := []int16{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}
	expected := []int16{0, 2, 2, 1, 1, 1, 2, 2, 2, -1, -1, -1, -1, -1, -1}
	compactData(&test)
	if !slices.Equal(test, expected) {
		t.Errorf(`Expected %v to match %v`, test, expected)
	}
}

func TestPart1(t *testing.T) {
	data := strings.NewReader(testData)
	result, err := Part1(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 1928 {
		t.Errorf("Expected result to be 1928. Got %d", result)
	}
}

func BenchmarkPart1(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		Part1(data)
	}
}
