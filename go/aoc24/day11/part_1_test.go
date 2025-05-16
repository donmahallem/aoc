package day11_test

import (
	"io"
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day11"
)

const testData string = "125 17"

func TestParseLine(t *testing.T) {
	test := "125 17"
	expected := []uint32{125, 17}
	reader := strings.NewReader(test)
	f1, _ := day11.ParseLine(reader)
	if !slices.Equal(f1, expected) {
		t.Errorf(`Expected %v to match %v`, f1, expected)
	}
}

func TestSplitStone(t *testing.T) {
	cache := make(map[day11.CacheKey]int)
	f1 := day11.SplitStone(125, 1, cache)
	if f1 != 1 {
		t.Errorf(`1Expected %d to match %d`, f1, 1)
	}
	f1 = day11.SplitStone(1256, 1, cache)
	if f1 != 2 {
		t.Errorf(`2Expected %d to match %d`, f1, 2)
	}
	f1 = day11.SplitStone(0, 1, cache)
	if f1 != 1 {
		t.Errorf(`3Expected %d to match %d`, f1, 1)
	}
}
func TestSplitStones(t *testing.T) {
	f1 := day11.SplitStones([]uint32{125, 17}, 6)
	if f1 != 22 {
		t.Errorf(`Expected %d to match %d`, f1, 22)
	}
	f1 = day11.SplitStones([]uint32{125, 17}, 25)
	if f1 != 55312 {
		t.Errorf(`Expected %d to match %d`, f1, 55312)
	}
}

func TestPart1(t *testing.T) {
	test := day11.Part1(strings.NewReader(testData))
	if test != 55312 {
		t.Errorf(`Expected %d to match 55312`, test)
	}
}

func BenchmarkParseLine(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day11.ParseLine(data)
	}
}

func BenchmarkPart1(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day11.Part1(data)
	}
}
