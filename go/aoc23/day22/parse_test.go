package day22

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed sample_1.txt
var testDataSample1 string

func Test_parseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		reader := strings.NewReader(testDataSample1)
		res, err := parseInput(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if len(res) != 7 {
			t.Errorf(`Expected number of bricks to be 7, got %d`, len(res))
		}

		expectedBricks := []brick{
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
