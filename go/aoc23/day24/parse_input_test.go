package day24

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

func Test_parseInput(t *testing.T) {
	t.Run("testData", func(t *testing.T) {
		reader := strings.NewReader(testData)
		res, err := parseInput[float64](reader)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBricks := []hail[float64]{
			{19, 13, 30, -2, 1, -2},
			{18, 19, 22, -1, -1, -2},
			{20, 25, 34, -2, -2, -4},
			{12, 31, 28, -1, -2, -1},
			{20, 19, 15, 1, -5, -3},
		}
		if len(res) != len(expectedBricks) {
			t.Errorf(`Expected number of bricks to be %d, got %d`, len(expectedBricks), len(res))
		}
		for i, hail := range res {
			if hail != expectedBricks[i] {
				t.Errorf("Expected hail %d to be %+v, got %+v", i, expectedBricks[i], hail)
			}
		}

	})

}

func Benchmark_parseInput(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {

		reader := strings.NewReader(testData)
		for b.Loop() {
			parseInput[float64](reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {

		full_data, ok := test_utils.GetTestData(23, 24)
		if ok {

			reader := strings.NewReader(full_data)
			for b.Loop() {
				parseInput[float64](reader)
				reader.Seek(0, 0)
			}
		} else {
			b.Skip("No full data available")
			return
		}
	})
}
