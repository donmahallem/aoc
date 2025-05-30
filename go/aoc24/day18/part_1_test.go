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

func TestParseInput(t *testing.T) {
	points := day18.ParseInput(strings.NewReader(testData))
	field := day18.PointsToField(points, 7, 7)
	if len(points) != 25 {
		t.Errorf(`Expected 25 items. Not %d`, points)
	}
	if len(field) != 7 {
		t.Errorf(`Expected %d to match 25`, len(field))
	}
}

func TestFindShortestPath(t *testing.T) {
	points := day18.ParseInput(strings.NewReader(testData))
	field := day18.PointsToField(points, 7, 7)
	if result := day18.FindShortestPath(field, 12, 7, 7); result != 22 {
		t.Errorf(`Expected %d to match 22`, result)
	}
}

func BenchmarkFindShortestPath(b *testing.B) {
	reader := strings.NewReader(testData)
	points := day18.ParseInput(reader)
	field := day18.PointsToField(points, 7, 7)
	for b.Loop() {
		day18.FindShortestPath(field, 12, 7, 7)
	}
}
func BenchmarkParseInput(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day18.ParseInput(reader)
	}
}
