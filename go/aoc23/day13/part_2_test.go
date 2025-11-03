package day13_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day13"
)

func TestOneBitApart(t *testing.T) {
	t.Run("one bit apart", func(t *testing.T) {
		a := 0b1100
		b := 0b1110
		if ok := day13.OneBitApart(a, b); !ok {
			t.Errorf(`Expected %04b and %04b to be one bit apart`, a, b)
		}
	})
	t.Run("not one bit apart", func(t *testing.T) {
		a := 0b1100
		b := 0b1010
		if ok := day13.OneBitApart(a, b); ok {
			t.Errorf(`Expected %04b and %04b to not be one bit apart`, a, b)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("test block 2", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day13.Part2(reader)
		if res != 400 {
			t.Errorf(`Expected number of blocks to be 400, got %d`, res)
		}
	})
}
