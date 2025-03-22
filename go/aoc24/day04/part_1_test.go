package day04_test

import (
	"fmt"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day04"
)

var testDataString = [][]byte{[]byte("MMMSXXMASM"),
	[]byte("MSAMXMSMSA"),
	[]byte("AMXSXMAAMM"),
	[]byte("MSAMASMSMX"),
	[]byte("XMASAMXAMM"),
	[]byte("XXAMMXXAMA"),
	[]byte("SMSMSASXSS"),
	[]byte("SAXAMASAAA"),
	[]byte("MAMMMXMMMM"),
	[]byte("MXMXAXMASX")}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func Test24Day04Part1(t *testing.T) {
	result := day04.CheckBlock(testDataString)
	if result != 18 {
		t.Errorf(`Expected %d to match %d`, result, 18)
	}
}

func Test24Day04Part1Slices(t *testing.T) {
	count := 0
	for i := 0; i < len(testDataString); i++ {
		count += day04.CheckBlock(testDataString[max(0, i-3) : i+1])
		fmt.Printf("%d - %d = %d\n", i-3, i+1, count)
	}
	if count != 18 {
		t.Errorf(`Expected %d to match %d`, count, 18)
	}
}
