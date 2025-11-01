package day22_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day22"
)

func TestPart2(t *testing.T) {
	t.Run("test sample 1", func(t *testing.T) {
		expected := 9
		reader := strings.NewReader(testDataSample1)
		result := day22.Part2(reader)
		if result != expected {
			t.Errorf(`Expected number of blocks to be %d, got %d`, expected, result)
		}
	})
}

func BenchmarkPart2(b *testing.B) {

	reader := strings.NewReader(testDataSample1)
	for b.Loop() {
		day22.Part2(reader)
		reader.Seek(0, 0)
	}
}
