package day04_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day04"
)

const testData string = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

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
func TestCheckBlock1(t *testing.T) {
	result := day04.CheckBlock(testDataString)
	if result != 18 {
		t.Errorf(`Expected %d to match %d`, result, 18)
	}
}

func TestCheckBlock2(t *testing.T) {
	count := 0
	for i := 0; i < len(testDataString); i++ {
		count += day04.CheckBlock(testDataString[max(0, i-3) : i+1])
	}
	if count != 30 {
		t.Errorf(`Expected %d to match %d`, count, 30)
	}
}

func Test24Day04Part1(t *testing.T) {

	if result := day04.Part1(strings.NewReader(testData)); result != 18 {
		t.Errorf(`Expected %d to match %d`, result, 18)
	}
}

func BenchmarkPart1(b *testing.B) {
	testData := strings.NewReader(testData)
	for b.Loop() {
		testData.Seek(0, io.SeekStart)
		day04.Part1(testData)
	}
}
