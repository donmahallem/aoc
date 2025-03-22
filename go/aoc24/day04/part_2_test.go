package day04_test

import (
	"testing"

	"github.com/donmahallem/aoc/aoc24/day04"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func Test24Day04Part2(t *testing.T) {
	result := day04.CheckMasBlock(testDataString)
	if result != 9 {
		t.Errorf(`Expected %d to match %d`, result, 9)
	}
}
