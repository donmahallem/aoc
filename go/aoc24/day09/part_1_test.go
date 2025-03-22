package day09_test

import (
	"slices"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day09"
)

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
