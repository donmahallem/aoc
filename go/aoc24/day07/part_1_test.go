package day07_test

import (
	"testing"

	"github.com/donmahallem/aoc/aoc24/day07"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName32(t *testing.T) {
	result := 190
	terms := []int{10, 19}
	if i := day07.CheckLinePart1(&result, &terms); !i {
		t.Errorf(`Expected %v to be %d and not %t`, terms, result, i)
	}
}
