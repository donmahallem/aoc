package day11_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day11"
)

func TestPart2(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 82000210
		reader := strings.NewReader(testData)
		if res := day11.Part2(reader); res != expected {
			t.Errorf(`Expected %d to be %d`, res, expected)
		}
	})
}
