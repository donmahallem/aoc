package day18_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day18"
)

const testData string = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

func TestFindShortestPath(t *testing.T) {
	parsedData := day18.ParseInput(strings.NewReader(testData), 7, 7)
	if result := day18.FindShortestPath(parsedData.Field, 12, 7, 7); result != 22 {
		t.Errorf(`Expected %d to match 22`, result)
	}
}

func BenchmarkFindShortestPath(b *testing.B) {
	reader := strings.NewReader(testData)
	parsedData := day18.ParseInput(reader, 7, 7)
	for b.Loop() {
		day18.FindShortestPath(parsedData.Field, 12, 7, 7)
	}
}

func TestPart1(t *testing.T) {
	if result := day18.Part1Base(strings.NewReader(testData), 12, 7, 7); result != 22 {
		t.Errorf(`Expected %d to match 22`, result)
	}
}

func BenchmarkParseInput(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day18.ParseInput(reader, 7, 7)
	}
}

func BenchmarkPart1Base(b *testing.B) {
	testData := strings.NewReader(testData)
	b.Run("sample dataset", func(b *testing.B) {
		for b.Loop() {
			testData.Seek(0, io.SeekStart)
			day18.Part1Base(testData, 12, 7, 7)
		}
	})
	b.Run("large dataset", func(b *testing.B) {
		for b.Loop() {
			testData.Seek(0, io.SeekStart)
			day18.Part1Base(testData, 5000, 71, 71)
		}
	})
}
