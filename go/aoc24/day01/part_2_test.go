package day01_test

import (
	"testing"

	"github.com/donmahallem/aoc/aoc24/day01"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func Test24Day04Part2(t *testing.T) {
	test := []int{1, 2, 4, 5, 5}
	result := day01.Count(test, 5)
	if result != 2 {
		t.Errorf(`Expected %d to match %d`, result, 2)
	}
}
