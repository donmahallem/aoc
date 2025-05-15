package day18_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day18"
	"github.com/donmahallem/aoc/test_utils"
)

func TestIsPathAvailable(t *testing.T) {
	reader := strings.NewReader(testData)
	points := day18.ParseInput(reader)
	field := day18.PointsToField(points, 7, 7)
	t.Run("test for 12 steps", func(t *testing.T) {
		ok := day18.IsPathAvailable(field, 5, 7, 7)
		if !ok {
			t.Errorf(`Expected to be true`)
		}
	})
	t.Run("test for 21 steps", func(t *testing.T) {
		ok := day18.IsPathAvailable(field, 5, 7, 7)
		if ok {
			t.Errorf(`Expected to be false`)
		}
	})
}

func TestFindFirstNonSolvable(t *testing.T) {
	reader := strings.NewReader(testData)
	points := day18.ParseInput(reader)
	field := day18.PointsToField(points, 7, 7)
	numPoints := uint16(len(points))
	t.Run("find an end", func(t *testing.T) {
		endPods := day18.FindFirstNonSolvable(field, numPoints, 7, 7)
		if endPods == numPoints {
			t.Errorf(`Expected result to be ok`)
		}
		expected := day18.Point{X: 6, Y: 1}
		if points[endPods] != expected {
			t.Errorf("Found %v. Expected: %v", points[endPods], expected)
		}
	})
	t.Run("find no end", func(t *testing.T) {
		endPods := day18.FindFirstNonSolvable(field, numPoints, 7, 7)
		if endPods == numPoints {
			t.Errorf(`Expected result to be false`)
		}
	})
}

func createEmptyField(width, height uint16) day18.Field {

	parsedField := make(day18.Field, height)
	for y := range height {
		parsedField[y] = make([]uint16, width)
	}
	return parsedField
}
func BenchmarkIsPathAvailable(b *testing.B) {
	const TEST_WIDTH, TEST_HEIGHT uint16 = 7, 7
	obstacleField := createEmptyField(TEST_WIDTH, TEST_HEIGHT)
	for b.Loop() {
		day18.IsPathAvailable(obstacleField, 5, TEST_WIDTH, TEST_HEIGHT)
	}
}
func BenchmarkFindFirstNonSolvable(b *testing.B) {
	b.Run("sample dataset", func(b *testing.B) {
		points := day18.ParseInput(strings.NewReader(testData))
		field := day18.PointsToField(points, 7, 7)
		for b.Loop() {
			day18.FindFirstNonSolvable(field, uint16(len(points)), 7, 7)
		}
	})
	b.Run("large dataset", func(b *testing.B) {
		if !test_utils.CheckTestDataExists(24, 18) {
			b.Skip("Couldn't retrieve test file data")
		}
		sourceData, _ := test_utils.GetTestData(24, 18)
		points := day18.ParseInput(strings.NewReader(sourceData))
		field := day18.PointsToField(points, 71, 71)
		for b.Loop() {
			day18.FindFirstNonSolvable(field, uint16(len(points)), 71, 71)
		}
	})
}
