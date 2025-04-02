package day06_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day06"
)

func TestOutOfBoundsShouldBeInside(t *testing.T) {
	x := int16(2)
	y := int16(2)
	testField := day06.NewField(10, 5, nil)
	if day06.OutOfBounds(&testField, &x, &y) {
		t.Errorf(`Expected %s to match`, "a")
	}
}

func TestOutOfBoundsShouldBeOutside(t *testing.T) {
	x := int16(-1)
	y := int16(5)
	testField := day06.NewField(6, 6, nil)
	if !day06.OutOfBounds(&testField, &x, &y) {
		t.Errorf(`Expected %s to match`, "a")
	}
	x = int16(1)
	y = int16(-5)
	testField = day06.NewField(6, 6, nil)
	if !day06.OutOfBounds(&testField, &x, &y) {
		t.Errorf(`Expected %s to match`, "a")
	}
}

func TestPart1(t *testing.T) {
	if res := day06.Part1(strings.NewReader(testData)); res != 41 {
		t.Errorf(`Expected %d to match %d`, res, 41)
	}
}
