package day19_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day19"
)

const testData string = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

const testPatterns string = "r, wr, b, g, bwu, rb, gb, br"

func TestParseFirstLine(t *testing.T) {
	testInput := testPatterns
	points, keyLen := day19.ParseFirstLine(&testInput)
	if len(*points) != 8 {
		t.Errorf(`Expected %d to match 8`, len(*points))
	}
	if *keyLen != 3 {
		t.Errorf(`Expected %d to match 3`, *keyLen)
	}
}
func TestParseInput(t *testing.T) {
	patterns, towls, keyLen := day19.ParseInput(strings.NewReader(testData))
	if len(*patterns) != 8 {
		t.Errorf(`Expected %d to match 8`, len(*patterns))
	}
	if len(*towls) != 8 {
		t.Errorf(`Expected %d to match 8`, len(*towls))
	}
	if *keyLen != 3 {
		t.Errorf(`Expected %d to match 3`, *keyLen)
	}
}
func TestPart1(t *testing.T) {
	test := day19.Part1(strings.NewReader(testData))
	if test != 6 {
		t.Errorf(`Expected %d to match 6`, test)
	}
}

func BenchmarkParseFirstLine(b *testing.B) {
	testInput := testPatterns
	for b.Loop() {
		day19.ParseFirstLine(&testInput)
	}
}

func BenchmarkParseInput(b *testing.B) {
	for b.Loop() {
		day19.ParseInput(strings.NewReader(testData))
	}
}

func BenchmarkPart1(b *testing.B) {
	for b.Loop() {
		day19.Part1(strings.NewReader(testData))
	}
}
