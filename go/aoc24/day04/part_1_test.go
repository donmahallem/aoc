package day04_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day04"
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

func Test24Day04Part1(t *testing.T) {

	result, err := day04.Part1(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 18 {
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
