package day18_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day18"
)

const testData = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

func TestPart1_Sample(t *testing.T) {
	t.Run("sample", func(t *testing.T) {
		result, err := day18.Part1Base(strings.NewReader(testData), 12, 7, 7)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if result != 22 {
			t.Errorf("Expected %v, got %v", 22, result)
		}
	})
}

func TestPart2_Sample(t *testing.T) {
	t.Run("sample", func(t *testing.T) {
		result, err := day18.Part2Base(strings.NewReader(testData), 7, 7)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		expected := []int16{6, 1}
		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
