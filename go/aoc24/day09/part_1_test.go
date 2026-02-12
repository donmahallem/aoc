package day09

import (
	"slices"
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
