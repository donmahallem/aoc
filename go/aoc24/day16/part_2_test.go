package day16_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day16"
)

func TestPart2(t *testing.T) {
	t.Run("with testData1", func(t *testing.T) {
		result := day16.Part2(strings.NewReader(testData1))
		if result != 45 {
			t.Errorf(`Expected %d to match 45`, result)
		}
	})
	t.Run("with testData2", func(t *testing.T) {
		result := day16.Part2(strings.NewReader(testData2))
		if result != 64 {
			t.Errorf(`Expected %d to match 64`, result)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	data := strings.NewReader(testData2)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day16.Part2(data)
	}
}
