package day04_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day04"
)

func Test24Day04Part2(t *testing.T) {

	if result := day04.Part2(strings.NewReader(testData)); result != 9 {
		t.Errorf(`Expected %d to match %d`, result, 9)
	}
}

func BenchmarkPart2(b *testing.B) {
	testData := strings.NewReader(testData)
	for b.Loop() {
		testData.Seek(0, io.SeekStart)
		day04.Part2(testData)
	}
}

func TestPart2NoMatches(t *testing.T) {
	const input = `XXX
XXX
XXX`
	if result := day04.Part2(strings.NewReader(input)); result != 0 {
		t.Errorf(`Expected %d to match %d`, result, 0)
	}
}

func TestPart2SingleCross(t *testing.T) {
	const input = `MXS
XAX
MXS`
	if result := day04.Part2(strings.NewReader(input)); result != 1 {
		t.Errorf(`Expected %d to match %d`, result, 1)
	}
}

func TestPart2MultipleCrosses(t *testing.T) {
	const input = `MXSXMXS
XAXXXAX
MXSXMXS
XAXXXAX
MXSXMXS`
	if result := day04.Part2(strings.NewReader(input)); result != 4 {
		t.Errorf(`Expected %d to match %d`, result, 4)
	}
}
