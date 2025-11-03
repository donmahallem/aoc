package day20_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day20"
)

func TestCountCheats2(t *testing.T) {
	patterns := day20.ParseInput(strings.NewReader(testData))
	if result := day20.CountCheats2(patterns, 50, 20); result != 285 {
		t.Errorf(`Expected %d to save %d steps. Not %d`, 2, 285, result)
	}
}

func BenchmarkCountCheats2(b *testing.B) {
	patterns := day20.ParseInput(strings.NewReader(testData))
	for b.Loop() {
		day20.CountCheats2(patterns, 50, 20)
	}
}
