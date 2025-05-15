package day08_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day08"
)

const testData string = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func Test24Day08Part1(t *testing.T) {

	if result := day08.Part1(strings.NewReader(testData)); result != 14 {
		t.Errorf(`Expected %d to match %d`, result, 14)
	}
}

func BenchmarkPart1(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day08.Part1(data)
	}
}
