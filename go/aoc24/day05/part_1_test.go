package day05_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day05"
)

const testData string = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestPart1(t *testing.T) {
	result := day05.Part1(strings.NewReader(testData))
	if result != 143 {
		t.Errorf(`Expected %d to be %d`, result, 143)
	}
}

func BenchmarkPart1(b *testing.B) {
	testData := strings.NewReader(testData)
	for b.Loop() {
		testData.Seek(0, io.SeekStart)
		day05.Part1(testData)
	}
}
