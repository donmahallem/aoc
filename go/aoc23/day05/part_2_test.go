package day05_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day05"
)

func TestPart2(t *testing.T) {
	result := day05.Part2(strings.NewReader(testData))
	if result != 46 {
		t.Errorf(`Expected %d to be %d`, result, 46)
	}
}
func TestIntersectInterval(t *testing.T) {
	t.Run("Non insterescting", func(t *testing.T) {
		a := day05.Interval{Start: 0, End: 100}
		b := day05.Interval{Start: 100, End: 200}
		result := day05.IntersectInterval(a, b)
		if result != nil {
			t.Errorf("Result should be nil")
		}
	})
	t.Run("Non insterescting", func(t *testing.T) {
		a := day05.Interval{Start: 0, End: 100}
		b := day05.Interval{Start: 50, End: 200}
		result := day05.IntersectInterval(a, b)
		expected := day05.Interval{Start: 50, End: 100}
		if result == nil {
			t.Errorf("Result should not be nil")
		} else if *result != expected {
			t.Errorf("Result should not be %v. Not %v", expected, *result)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day05.Part2(reader)
	}
}
