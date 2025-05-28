package day09_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day09"
)

func TestPredictLeft(t *testing.T) {
	t.Run("row_1", func(t *testing.T) {
		testRow := []int{0, 3, 6, 9, 12, 15}
		result := day09.PredictLeft(testRow)
		if result != -3 {
			t.Errorf(`Expected result to be -3 not %d`, result)
		}
	})
	t.Run("row_2", func(t *testing.T) {
		testRow := []int{1, 3, 6, 10, 15, 21}
		result := day09.PredictLeft(testRow)
		if result != 0 {
			t.Errorf(`Expected result to be 0 not %d`, result)
		}
	})
	t.Run("row_3", func(t *testing.T) {
		testRow := []int{10, 13, 16, 21, 30, 45}
		result := day09.PredictLeft(testRow)
		if result != 5 {
			t.Errorf(`Expected result to be 5 not %d`, result)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 2
		reader := strings.NewReader(testData)
		if res := day09.Part2(reader); res != expected {
			t.Errorf(`Expected %d to be %d`, res, expected)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day09.Part2(reader)
	}
}
