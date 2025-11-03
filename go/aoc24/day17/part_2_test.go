package day17_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day17"
)

const testData2 string = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

func TestPart2_testData2(t *testing.T) {
	result := day17.Part2(strings.NewReader(testData2))
	if result != 117440 {
		t.Errorf(`Expected %d to match 117440`, result)
	}
}

func BenchmarkPart2(b *testing.B) {
	data := strings.NewReader(testData2)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day17.Part2(data)
	}
}
