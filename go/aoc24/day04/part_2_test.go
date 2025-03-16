package day04

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func Test24Day04Part2(t *testing.T) {
	result := CheckMasBlock(testDataString)
	if result != 9 {
		t.Errorf(`Expected %d to match %d`, result, 9)
	}
}
