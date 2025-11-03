package day07_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day07"
)

const testData string = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

const testData2 string = `32T3K 100
43T4Q 200`

const testData3 string = `32T3K 100
43T4Q 200
AQKQA 150
QKQAK 250
2777K 50
2888A 75
AAAAK 300
KKKKA 350
KQQKQ 250`

func TestPart1(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		var expected int = 6440
		reader := strings.NewReader(testData)
		if res := day07.Part1(reader); res != expected {
			t.Errorf(`Expected %v to match %v`, res, expected)
		}
	})
	t.Run("testData2", func(t *testing.T) {
		var expected int = 500
		reader := strings.NewReader(testData2)
		if res := day07.Part1(reader); res != expected {
			t.Errorf(`Expected %v to match %v`, res, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day07.Part1(reader)
	}
}
