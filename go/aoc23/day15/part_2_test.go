package day15_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day15"
)

func TestParseInput(t *testing.T) {
	t.Run("test parse input", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day15.ParseInput(reader)
		expectedLen := 11
		if len(res) != expectedLen {
			t.Errorf(`Expected number of groups to be %d, got %d`, expectedLen, len(res))
		}
	})
}

func TestPart2(t *testing.T) {

	t.Run("test sample 2", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day15.Part2(reader)
		if res != 145 {
			t.Errorf(`Expected number of blocks to be 405, got %d`, res)
		}
	})
}
