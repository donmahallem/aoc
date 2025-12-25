package day13

import (
	"testing"
)

func TestOneBitApart(t *testing.T) {
	t.Run("one bit apart", func(t *testing.T) {
		a := 0b1100
		b := 0b1110
		if ok := oneBitApart(a, b); !ok {
			t.Errorf(`Expected %04b and %04b to be one bit apart`, a, b)
		}
	})
	t.Run("not one bit apart", func(t *testing.T) {
		a := 0b1100
		b := 0b1010
		if ok := oneBitApart(a, b); ok {
			t.Errorf(`Expected %04b and %04b to not be one bit apart`, a, b)
		}
	})
}
