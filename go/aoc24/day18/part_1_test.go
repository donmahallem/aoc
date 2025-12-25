package day18_test

import (
	_ "embed"
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day18"
)

//go:embed sample.txt
var testData string

func TestFindShortestPath(t *testing.T) {
	parsedData, _ := day18.ParseInput(strings.NewReader(testData), 7, 7)
	if result := day18.FindShortestPath(parsedData.Field, 12, 7, 7); result != 22 {
		t.Errorf(`Expected %d to match 22`, result)
	}
}

func BenchmarkFindShortestPath(b *testing.B) {
	reader := strings.NewReader(testData)
	parsedData, _ := day18.ParseInput(reader, 7, 7)
	for b.Loop() {
		day18.FindShortestPath(parsedData.Field, 12, 7, 7)
	}
}

func TestPart1(t *testing.T) {
	if result, _ := day18.Part1Base(strings.NewReader(testData), 12, 7, 7); result != 22 {
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
