package day17_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day17"
)

var testData string = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

func TestParseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 13 * 13
		reader := strings.NewReader(testData)
		if res := day17.ParseInput(reader); len(res.Cells) != expected {
			t.Errorf(`Expected %d to be %d`, len(res.Cells), expected)
		}

	})

}

func TestPart1(t *testing.T) {
	t.Run("test block 2", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day17.Part1(reader)
		if res != 102 {
			t.Errorf(`Expected number of blocks to be 102, got %d`, res)
		}
	})
}

func BenchmarkPart1(b *testing.B) {

	reader := strings.NewReader(testData)
	for b.Loop() {
		day17.Part1(reader)
		reader.Seek(0, 0)
	}
}
