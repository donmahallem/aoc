package day21_test

import (
	"fmt"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day21"
)

func TestWalkNumericSequence(t *testing.T) {
	var testSequence []byte = []byte{'0', '2', '9', 'A'}
	t.Run(fmt.Sprintf("%v width depth 1", string(testSequence)), func(t *testing.T) {
		// <A^A>^^AvvvA
		cache := make(day21.Cache)
		result := day21.WalkNumericSequence(&testSequence, 1, &cache)
		if result != 12 {
			t.Errorf(`Expected %d to match 12`, result)
		}
	})
	t.Run(fmt.Sprintf("%v width depth 2", string(testSequence)), func(t *testing.T) {
		// v<<A>>^A<A>AvA<^AA>A<vAAA>^A
		cache := make(day21.Cache)
		result := day21.WalkNumericSequence(&testSequence, 2, &cache)
		if result != 28 {
			t.Errorf(`Expected %d to match 28`, result)
		}
	})
	t.Run(fmt.Sprintf("%v width depth 3", string(testSequence)), func(t *testing.T) {
		// <vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A
		cache := make(day21.Cache)
		result := day21.WalkNumericSequence(&testSequence, 3, &cache)
		if result != 68 {
			t.Errorf(`Expected %d to match 68`, result)
		}
	})
}
