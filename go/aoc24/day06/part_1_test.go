package day06

import (
	"testing"
)

func TestOutOfBoundsShouldBeInside(t *testing.T) {
	x := int16(2)
	y := int16(2)
	testField := NewField(10, 5, nil)
	if OutOfBounds(&testField, &x, &y) {
		t.Errorf(`Expected %s to match`, "a")
	}
}

func TestOutOfBoundsShouldBeOutside(t *testing.T) {
	x := int16(-1)
	y := int16(5)
	testField := NewField(6, 6, nil)
	if !OutOfBounds(&testField, &x, &y) {
		t.Errorf(`Expected %s to match`, "a")
	}
	x = int16(1)
	y = int16(-5)
	testField = NewField(6, 6, nil)
	if !OutOfBounds(&testField, &x, &y) {
		t.Errorf(`Expected %s to match`, "a")
	}
}
