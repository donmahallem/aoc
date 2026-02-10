package day09

import (
	"testing"
)

func TestPredictLeft(t *testing.T) {
	t.Run("row_1", func(t *testing.T) {
		testRow := []int{0, 3, 6, 9, 12, 15}
		result := predictLeft(testRow)
		if result != -3 {
			t.Errorf(`Expected result to be -3 not %d`, result)
		}
	})
	t.Run("row_2", func(t *testing.T) {
		testRow := []int{1, 3, 6, 10, 15, 21}
		result := predictLeft(testRow)
		if result != 0 {
			t.Errorf(`Expected result to be 0 not %d`, result)
		}
	})
	t.Run("row_3", func(t *testing.T) {
		testRow := []int{10, 13, 16, 21, 30, 45}
		result := predictLeft(testRow)
		if result != 5 {
			t.Errorf(`Expected result to be 5 not %d`, result)
		}
	})
}
