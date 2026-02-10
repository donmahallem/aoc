package day20

import (
	"strings"
	"testing"
)

func TestCountCheats2(t *testing.T) {
	patterns, err := parseInput(strings.NewReader(testData))
	if err != nil {
		t.Errorf(`Unexpected error: %v`, err)
	}
	if result := CountCheats2(patterns, 50, 20); result != 285 {
		t.Errorf(`Expected %d to save %d steps. Not %d`, 2, 285, result)
	}
}

func BenchmarkCountCheats2(b *testing.B) {
	patterns, err := parseInput(strings.NewReader(testData))
	if err != nil {
		b.Errorf(`Unexpected error: %v`, err)
	}
	for b.Loop() {
		CountCheats2(patterns, 50, 20)
	}
}

func FuzzPart2(f *testing.F) {
	f.Add(testData)
	f.Fuzz(func(t *testing.T, input string) {
		r := strings.NewReader(input)
		_, err := Part2(r)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}
