package day10

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/test_utils"
)

const testData string = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestWalkDepth(t *testing.T) {
	f1 := loadField(strings.NewReader(testData))

	result := make(map[uint16]struct{})
	walkDepth(f1, 2, 0, result)
	if len(result) != 5 {
		t.Errorf(`Expected %v to not match %d`, len(result), 5)
	}
	result = make(map[uint16]struct{})
	walkDepth(f1, 4, 0, result)
	if len(result) != 6 {
		t.Errorf(`Expected %v to not match %d`, len(result), 5)
	}
}
func TestSearchAll(t *testing.T) {
	f1 := loadField(strings.NewReader(testData))
	result := searchAll(f1)
	if result != 36 {
		t.Errorf(`Expected %d to not match %d`, result, 36)
	}
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		data := strings.NewReader(testData)
		for b.Loop() {
			data.Seek(0, io.SeekStart)
			Part1(data)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 24, 10, Part1)
	})
}
