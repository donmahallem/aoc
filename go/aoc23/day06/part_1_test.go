package day06_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day06"
)

const testData string = `Time:      7  15   30
Distance:  9  40  200`

func TestPart1(t *testing.T) {
	result := day06.Part1(strings.NewReader(testData))
	if result != 288 {
		t.Errorf(`Expected %d to be %d`, result, 288)
	}
}
func TestCountOptions(t *testing.T) {
	t.Run("7_9", func(t *testing.T) {
		testItem := day06.Race{Time: 7, Distance: 9}
		result := day06.CountOptions(testItem)
		if result != 4 {
			t.Errorf(`Expected %d to be %d`, result, 4)
		}
	})
	t.Run("15_40", func(t *testing.T) {
		testItem := day06.Race{Time: 15, Distance: 40}
		result := day06.CountOptions(testItem)
		if result != 8 {
			t.Errorf(`Expected %d to be %d`, result, 8)
		}
	})
	t.Run("30_200", func(t *testing.T) {
		testItem := day06.Race{Time: 30, Distance: 200}
		result := day06.CountOptions(testItem)
		if result != 9 {
			t.Errorf(`Expected %d to be %d`, result, 9)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day06.Part1(reader)
	}
}
