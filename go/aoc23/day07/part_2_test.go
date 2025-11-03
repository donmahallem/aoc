package day07_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day07"
)

func TestPart2(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		var expected int = 5905
		reader := strings.NewReader(testData)
		if res := day07.Part2(reader); res != expected {
			t.Errorf(`Expected %v to match %v`, res, expected)
		}
	})
	t.Run("testData3", func(t *testing.T) {
		var expected int = 9800
		reader := strings.NewReader(testData3)
		if res := day07.Part2(reader); res != expected {
			t.Errorf(`Expected %v to match %v`, res, expected)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day07.Part2(reader)
	}
}
