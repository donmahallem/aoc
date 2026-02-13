package day11

import (
	"io"
	"slices"
	"strings"
	"testing"
)

const testData string = "125 17"

func TestParseLine(t *testing.T) {
	test := "125 17"
	expected := []int{125, 17}
	reader := strings.NewReader(test)
	f1, _ := parseLine(reader)
	if !slices.Equal(f1, expected) {
		t.Errorf(`Expected %v to match %v`, f1, expected)
	}
}

func TestSplitStone(t *testing.T) {
	cache := make(map[cacheKey]int)
	f1 := splitStone(125, 1, cache)
	if f1 != 1 {
		t.Errorf(`1Expected %d to match %d`, f1, 1)
	}
	f1 = splitStone(1256, 1, cache)
	if f1 != 2 {
		t.Errorf(`2Expected %d to match %d`, f1, 2)
	}
	f1 = splitStone(0, 1, cache)
	if f1 != 1 {
		t.Errorf(`3Expected %d to match %d`, f1, 1)
	}
}
func TestSplitStones(t *testing.T) {
	f1 := splitStones([]int{125, 17}, 6)
	if f1 != 22 {
		t.Errorf(`Expected %d to match %d`, f1, 22)
	}
	f1 = splitStones([]int{125, 17}, 25)
	if f1 != 55312 {
		t.Errorf(`Expected %d to match %d`, f1, 55312)
	}
}

func BenchmarkParseLine(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		parseLine(data)
	}
}
