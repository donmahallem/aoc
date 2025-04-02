package day04_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day04"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestCheckMasBlock(t *testing.T) {
	result := day04.CheckMasBlock(testDataString)
	if result != 9 {
		t.Errorf(`Expected %d to match %d`, result, 9)
	}
}

func Test24Day04Part2(t *testing.T) {

	if result := day04.Part2(strings.NewReader(testData)); result != 9 {
		t.Errorf(`Expected %d to match %d`, result, 9)
	}
}
