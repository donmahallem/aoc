package day18_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day18"
)

var testData string = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

func TestParseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 14
		reader := strings.NewReader(testData)
		if res := day18.ParseInput(reader, true); len(res) != expected {
			t.Errorf(`Expected %d to be %d`, len(res), expected)
		}

	})

}

func TestPart1(t *testing.T) {
	t.Run("test block 2", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day18.Part1(reader)
		if res != 62 {
			t.Errorf(`Expected number of blocks to be 62, got %d`, res)
		}
	})
}

func BenchmarkPart1(b *testing.B) {

	reader := strings.NewReader(testData)
	for b.Loop() {
		day18.Part1(reader)
		reader.Seek(0, 0)
	}
}
