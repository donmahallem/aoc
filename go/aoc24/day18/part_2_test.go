package day18_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day18"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestIsPathAvailable(t *testing.T) {
	reader := strings.NewReader(testData)
	parsedData := day18.ParseInput(reader, 7, 7)
	t.Run("test for 12 steps", func(t *testing.T) {
		ok := day18.IsPathAvailable(parsedData.Field, 5, 7, 7)
		if !ok {
			t.Errorf(`Expected to be true`)
		}
	})
	t.Run("test for 21 steps", func(t *testing.T) {
		ok := day18.IsPathAvailable(parsedData.Field, 25, 7, 7)
		if ok {
			t.Errorf(`Expected to be false`)
		}
	})
}

func BenchmarkIsPathAvailable(b *testing.B) {
	const TEST_WIDTH, TEST_HEIGHT int16 = 7, 7
	obstacleField := make(day18.Field, TEST_HEIGHT*TEST_WIDTH)
	for b.Loop() {
		day18.IsPathAvailable(obstacleField, 5, TEST_WIDTH, TEST_HEIGHT)
	}
}
func BenchmarkFindFirstNonSolvable(b *testing.B) {
	b.Run("sample dataset", func(b *testing.B) {
		parsedData := day18.ParseInput(strings.NewReader(testData), 7, 7)
		for b.Loop() {
			day18.FindFirstNonSolvable(parsedData.Field, int16(len(parsedData.CorruptionOrder)), 7, 7)
		}
	})
	b.Run("large dataset", func(b *testing.B) {
		if !test_utils.CheckTestDataExists(24, 18) {
			b.Skip("Couldn't retrieve test file data")
		}
		sourceData, _ := test_utils.GetTestData(24, 18)
		parsedData := day18.ParseInput(strings.NewReader(sourceData), 71, 71)
		for b.Loop() {
			day18.FindFirstNonSolvable(parsedData.Field, int16(len(parsedData.CorruptionOrder)), 71, 71)
		}
	})
}

func TestPart2(t *testing.T) {
	if result := day18.Part2Base(strings.NewReader(testData), 7, 7); result != (day18.Point{X: 6, Y: 1}) {
		t.Errorf(`Expected %v to contain %v`, result, day18.Point{X: 6, Y: 1})
	}
}
