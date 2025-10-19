package day12_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day12"
)

var testData string = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func TestParseInput(t *testing.T) {
	t.Run("testData", func(t *testing.T) {
		reader := strings.NewReader(testData)
		lines := day12.ParseInput(reader, 1)
		const expectedLen int = 6
		if len(lines) != expectedLen {
			t.Errorf(`Expected %d lines to be parsed, got %d`, expectedLen, len(lines))
		}
	})
}

func TestPart1(t *testing.T) {
	t.Run("testData", func(t *testing.T) {
		reader := strings.NewReader(testData)
		result := day12.Part1(reader)
		const expectedLen int = 21
		if result != expectedLen {
			t.Errorf(`Expected %d galaxies to be parsed, got %d`, expectedLen, result)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day12.Part1(reader)
	}
}
