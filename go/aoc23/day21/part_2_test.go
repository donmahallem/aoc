package day21_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day21"
)

func TestPart2(t *testing.T) {
	var testData = []struct {
		input    int
		expected int
	}{
		{6, 16},
		{10, 50},
		{50, 1594},
		{100, 6536},
		{500, 167004},
		{1000, 668697},
		//{5000, 16733044}, takes too long...
	}

	for _, td := range testData {
		t.Run(fmt.Sprintf("test sample %d", td.input), func(t *testing.T) {
			reader := strings.NewReader(testDataSample1)
			res := day21.ParseInput(reader)
			visitedCount := day21.CountVisited(&res, td.input, false)
			if visitedCount != td.expected {
				t.Errorf(`Expected number of blocks to be %d, got %d`, td.expected, visitedCount)
			}
		})
	}
}

func BenchmarkPart2(b *testing.B) {

	reader := strings.NewReader(testDataSample1)
	for b.Loop() {
		day21.Part2(reader)
		reader.Seek(0, 0)
	}
}
