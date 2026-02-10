package day07

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

func Test_parseInput(t *testing.T) {
	reader := strings.NewReader(testData)
	inp, err := parseInput(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	x, y := inp.startX, inp.startY
	s := inp.startNode
	w, h := inp.width, inp.height
	if x != 7 || y != 0 {
		t.Errorf("unexpected start position: got (%d,%d), want (7,0)", x, y)
	}
	if s == nil {
		t.Errorf("unexpected start node: got nil, want non-nil")
	}
	if w != 15 {
		t.Errorf("unexpected width: got %d, want 15", w)
	}
	if h != 16 {
		t.Errorf("unexpected height: got %d, want 10", h)
	}
}

func Benchmark_parseInput(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {

		reader := strings.NewReader(testData)
		for b.Loop() {
			parseInput(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		fullTestData, ok := test_utils.GetTestData(25, 7)
		if !ok {
			b.Skipf("Test data for %d day %d not found, skipping...", 25, 7)
		}
		reader := strings.NewReader(fullTestData)
		for b.Loop() {
			parseInput(reader)
			reader.Seek(0, 0)
		}
	})
}
