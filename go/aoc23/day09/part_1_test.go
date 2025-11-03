package day09_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day09"
)

var testData string = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestParseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		reader := strings.NewReader(testData)
		res := day09.ParseInput(reader)
		if len(res) != 3 {
			t.Errorf(`Expected 3 rows not %d`, len(res))
		}
		if len(res[0]) != 6 {
			t.Errorf(`Expected 6 numbers not %d`, res[0])
		}
	})
}

func TestPredictRight(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		testRow := []int{0, 3, 6, 9, 12, 15}
		result := day09.PredictRight(testRow)
		if result != 18 {
			t.Errorf(`Expected 3 rows not %d`, result)
		}
	})
}

func TestPart1(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 114
		reader := strings.NewReader(testData)
		if res := day09.Part1(reader); res != expected {
			t.Errorf(`Expected %d to be %d`, res, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day09.Part1(reader)
	}
}
