package day23

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed sample.txt
var testData string

func Test_countDegrees(t *testing.T) {
	t.Run("respect slope", func(t *testing.T) {
		t.Run("test sample", func(t *testing.T) {
			reader := strings.NewReader(testData)
			result, w, h := parseInput(reader, true)
			expectedDegrees := map[[2]int][2]int{
				{1, 0}: {1, 1},
				{1, 1}: {2, 2},
				{3, 4}: {1, 2},
				{3, 5}: {1, 2},
			}
			for idx, expectedVal := range expectedDegrees {
				pos := idx[1]*w + idx[0]
				entries, exits := countDegrees(result, pos, w, h)
				if entries != expectedVal[0] || exits != expectedVal[1] {
					t.Errorf(`Expected index %d (%d,%d) to have %d entries and %d exits, got %d and %d`, pos, idx[0], idx[1], expectedVal[0], expectedVal[1], entries, exits)
				}
			}
		})
	})
}
