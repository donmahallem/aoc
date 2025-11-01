package day22_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day22"
)

var testDataSample1 string = `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

func TestParseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		reader := strings.NewReader(testDataSample1)
		res := day22.ParseInput(reader)
		if len(res) != 7 {
			t.Errorf(`Expected number of bricks to be 7, got %d`, len(res))
		}

		expectedBricks := []day22.Brick{
			{X1: 1, Y1: 0, Z1: 1, X2: 1, Y2: 2, Z2: 1},
			{X1: 0, Y1: 0, Z1: 2, X2: 2, Y2: 0, Z2: 2},
			{X1: 0, Y1: 2, Z1: 3, X2: 2, Y2: 2, Z2: 3},
			{X1: 0, Y1: 0, Z1: 4, X2: 0, Y2: 2, Z2: 4},
			{X1: 2, Y1: 0, Z1: 5, X2: 2, Y2: 2, Z2: 5},
			{X1: 0, Y1: 1, Z1: 6, X2: 2, Y2: 1, Z2: 6},
			{X1: 1, Y1: 1, Z1: 8, X2: 1, Y2: 1, Z2: 9},
		}
		for i, brick := range res {
			if brick != expectedBricks[i] {
				t.Errorf(`Expected brick %d to be %+v, got %+v`, i, expectedBricks[i], brick)
			}
		}

	})

}

func TestPart1(t *testing.T) {
	t.Run("test sample 1", func(t *testing.T) {
		expected := 5
		reader := strings.NewReader(testDataSample1)
		result := day22.Part1(reader)
		if result != expected {
			t.Errorf(`Expected number of blocks to be %d, got %d`, expected, result)
		}
	})
}

func BenchmarkPart1(b *testing.B) {

	reader := strings.NewReader(testDataSample1)
	for b.Loop() {
		day22.Part1(reader)
		reader.Seek(0, 0)
	}
}
