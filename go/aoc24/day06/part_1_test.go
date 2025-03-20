package day06

import (
	"testing"
)

func TestOutOfBoundsShouldBeInside(t *testing.T) {
	x := 2
	y := 2
	testField := NewField(10, 5, nil)
	if OutOfBounds(&testField, &x, &y) {
		t.Errorf(`Expected %s to match`, "a")
	}
}

func TestOutOfBoundsShouldBeOutside(t *testing.T) {
	x := -1
	y := 5
	testField := NewField(6, 6, nil)
	if !OutOfBounds(&testField, &x, &y) {
		t.Errorf(`Expected %s to match`, "a")
	}
	x = 1
	y = -5
	testField = NewField(6, 6, nil)
	if !OutOfBounds(&testField, &x, &y) {
		t.Errorf(`Expected %s to match`, "a")
	}
}
