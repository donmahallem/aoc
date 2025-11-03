package day02_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day02"
)

func TestCalculateMinBlock(t *testing.T) {
	data := []byte(testDataSlices[0])
	_, blocks := day02.ParseLine(data)
	if res := day02.CalculateMinBlock(&blocks); res != 48 {
		t.Errorf(`Expected %d to be %d`, res, 48)
	}
}

func TestPart2(t *testing.T) {
	result := day02.Part2(strings.NewReader(testData))
	if result != 2286 {
		t.Errorf(`Expected %d to be %d`, result, 2286)
	}
}

func BenchmarkPart2(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day02.Part2(reader)
	}
}
