package day06

import (
	"testing"
)

func TestCountOptions(t *testing.T) {
	t.Run("7_9", func(t *testing.T) {
		testItem := race{Time: 7, Distance: 9}
		result := countOptions(testItem)
		if result != 4 {
			t.Errorf(`Expected %d to be %d`, result, 4)
		}
	})
	t.Run("15_40", func(t *testing.T) {
		testItem := race{Time: 15, Distance: 40}
		result := countOptions(testItem)
		if result != 8 {
			t.Errorf(`Expected %d to be %d`, result, 8)
		}
	})
	t.Run("30_200", func(t *testing.T) {
		testItem := race{Time: 30, Distance: 200}
		result := countOptions(testItem)
		if result != 9 {
			t.Errorf(`Expected %d to be %d`, result, 9)
		}
	})
}
