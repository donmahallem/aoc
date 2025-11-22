package day24

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/test_utils"
)

func Test_findCollisions(t *testing.T) {
	t.Run("testData", func(t *testing.T) {
		reader := strings.NewReader(testData)
		inp := parseInput[float64](reader)
		collisions := findCollisions(inp, 7, 27)
		if collisions != 2 {
			t.Errorf("Expected %d to be 2", collisions)
		}

	})

	t.Run("run on full test data", func(t *testing.T) {
		full_data, ok := test_utils.GetTestData(23, 24)
		if !ok {
			t.Skip("test data not available")
			return
		}
		reader := strings.NewReader(full_data)
		inp := parseInput[float64](reader)
		collisions := findCollisions(inp, 200000000000000, 400000000000000)
		if collisions != 2 {
			t.Errorf("Expected %d to be 2", collisions)
		}
	})

}

func Benchmark_findCollisions(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {

		reader := strings.NewReader(testData)
		inp := parseInput[float64](reader)
		for b.Loop() {
			findCollisions(inp, 7, 27)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {

		full_data, ok := test_utils.GetTestData(23, 24)
		if ok {
			reader := strings.NewReader(full_data)
			inp := parseInput[float64](reader)

			for b.Loop() {
				findCollisions(inp, 200000000000000, 400000000000000)
			}
		} else {
			b.Skip("No full data available")
			return
		}
	})
}
