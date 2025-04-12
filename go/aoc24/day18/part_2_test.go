package day18_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day18"
	"github.com/donmahallem/aoc/test_utils"
)

func TestIsPathAvailable(t *testing.T) {
	points := day18.ParseInput(strings.NewReader(testData))
	t.Run("test for 12 steps", func(t *testing.T) {
		field := day18.ConvertInputToField(points, 12, 7, 7)
		scoreField := day18.CreateEmptyField(7, 7)
		ok := day18.IsPathAvailable(field, scoreField, 5, 7, 7)
		if !ok {
			t.Errorf(`Expected to be true`)
		}
	})
	t.Run("test for 21 steps", func(t *testing.T) {
		field := day18.ConvertInputToField(points, 22, 7, 7)
		scoreField := day18.CreateEmptyField(7, 7)
		ok := day18.IsPathAvailable(field, scoreField, 5, 7, 7)
		if ok {
			t.Errorf(`Expected to be false`)
		}
	})
}

func BenchmarkFindFirstNonSolvable(b *testing.B) {
	b.Run("sample dataset", func(b *testing.B) {
		points := day18.ParseInput(strings.NewReader(testData))
		for b.Loop() {
			day18.FindFirstNonSolvable(points, 71, 71)
		}
	})
	b.Run("large dataset", func(b *testing.B) {
		if !test_utils.CheckTestDataExists(24, 18) {
			b.Skip("Couldn't retrieve test file data")
		}
		sourceData, _ := test_utils.GetTestData(24, 18)
		points := day18.ParseInput(strings.NewReader(sourceData))
		for b.Loop() {
			day18.FindFirstNonSolvable(points, 71, 71)
		}
	})
}
