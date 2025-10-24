package day16_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day16"
)

var testData string = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func TestParseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 10
		reader := strings.NewReader(testData)
		if res := day16.ParseInputPart1(reader); len(res.Cells) != expected {
			t.Errorf(`Expected %d to be %d`, len(res.Cells), expected)
		}

	})

}

func TestPart1(t *testing.T) {
	t.Run("test block 2", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day16.Part1(reader)
		if res != 46 {
			t.Errorf(`Expected number of blocks to be 46, got %d`, res)
		}
	})
}

func BenchmarkPart1(b *testing.B) {

	reader := strings.NewReader(testData)
	for b.Loop() {
		day16.Part1(reader)
		reader.Seek(0, 0)
	}
}
