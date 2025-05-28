package day08_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day08"
)

var testData3 string = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestPart2(t *testing.T) {
	t.Run("testData3", func(t *testing.T) {
		const expected uint = 6
		reader := strings.NewReader(testData3)
		if res := day08.Part2(reader); res != expected {
			t.Errorf(`Expected %d to be %d`, res, expected)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day08.Part2(reader)
	}
}
