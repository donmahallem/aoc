package day10_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day10"
)

const testData string = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestFieldCompare(t *testing.T) {
	f1_1 := [][]byte{{5, 5}}
	start_1 := []day10.Position{{X: 1, Y: 2}}
	f1 := day10.NewField(2, 2, f1_1, start_1)

	f2_1 := [][]byte{{5, 5}}
	start_2 := []day10.Position{{X: 1, Y: 2}}
	f2 := day10.NewField(2, 2, f2_1, start_2)

	f3_1 := [][]byte{{5, 5}}
	start_3 := []day10.Position{{X: 1, Y: 2}, {X: 3, Y: 2}}
	f3 := day10.NewField(2, 2, f3_1, start_3)
	if !f1.Compare(&f2) {
		t.Errorf(`Expected %v to match %v`, f2, f1)
	}
	if f1.Compare(&f3) {
		t.Errorf(`Expected %v to not match %v`, f3, f1)
	}
}

func TestLoadField(t *testing.T) {
	f1, _ := day10.LoadField(strings.NewReader(testData))
	testSubField := [][]byte{[]uint8{8, 9, 0, 1, 0, 1, 2, 3},
		[]uint8{7, 8, 1, 2, 1, 8, 7, 4},
		[]uint8{8, 7, 4, 3, 0, 9, 6, 5},
		[]uint8{9, 6, 5, 4, 9, 8, 7, 4},
		[]uint8{4, 5, 6, 7, 8, 9, 0, 3},
		[]uint8{3, 2, 0, 1, 9, 0, 1, 2},
		[]uint8{0, 1, 3, 2, 9, 8, 0, 1},
		[]uint8{1, 0, 4, 5, 6, 7, 3, 2}}
	testStarts := []day10.Position{{X: 2, Y: 0},
		{X: 4, Y: 0},
		{X: 4, Y: 2},
		{X: 6, Y: 4},
		{X: 2, Y: 5},
		{X: 5, Y: 5},
		{X: 0, Y: 6},
		{X: 6, Y: 6},
		{X: 1, Y: 7}}
	testField := day10.NewField(8, 8, testSubField, testStarts)
	if !f1.Compare(&testField) {
		t.Errorf(`Expected %v to not match %v`, f1, testField)
	}
}

func TestWalkDepth(t *testing.T) {
	f1, _ := day10.LoadField(strings.NewReader(testData))

	result := make(map[day10.Position]bool)
	day10.WalkDepth(&f1, 2, 0, 0, &result)
	if len(result) != 5 {
		t.Errorf(`Expected %v to not match %d`, len(result), 5)
	}
	result = make(map[day10.Position]bool)
	day10.WalkDepth(&f1, 4, 0, 0, &result)
	if len(result) != 6 {
		t.Errorf(`Expected %v to not match %d`, len(result), 5)
	}
}
func TestSearchAll(t *testing.T) {
	f1, _ := day10.LoadField(strings.NewReader(testData))
	result := day10.SearchAll(&f1)
	if result != 36 {
		t.Errorf(`Expected %d to not match %d`, result, 36)
	}
}

func TestPart1(t *testing.T) {
	if result := day10.Part1(strings.NewReader(testData)); result != 36 {
		t.Errorf(`Expected %d to contain %d`, result, 36)
	}
}

func BenchmarkPart1(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day10.Part1(data)
	}
}
