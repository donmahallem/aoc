package day23_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day23"
)

func TestStringifySequence(t *testing.T) {
	seq := day23.NodeHashList{day23.HashId([]byte{'a', 'b'}), day23.HashId([]byte{'z', 'u'})}
	expected := "ab,zu"
	if res := day23.StringifySequence(seq); res != expected {
		t.Errorf("String doesn't match. Expect '%s' not '%s'", expected, res)
	}
}

func TestFindLongest(t *testing.T) {
	points := day23.ParseInputMap(strings.NewReader(testData))
	data := day23.FindLongest(points)
	if len(data) != 4 {
		t.Errorf(`Expected %d to match 8`, data)
	}
}

func TestPart2(t *testing.T) {
	data := day23.Part2(strings.NewReader(testData))
	expected := "co,de,ka,ta"
	if data != expected {
		t.Errorf(`Expected %s to match %s`, data, expected)
	}
}

func BenchmarkFindLongest(b *testing.B) {
	points := day23.ParseInputMap(strings.NewReader(testData))
	for b.Loop() {
		day23.FindLongest(points)
	}
}
func BenchmarkPart2(b *testing.B) {
	for b.Loop() {
		day23.Part2(strings.NewReader(testData))
	}
}
