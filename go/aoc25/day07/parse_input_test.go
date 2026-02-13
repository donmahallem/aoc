package day07

import (
	"strings"
	"testing"
)

const testData = "..^..\n..S..\n..^..\n..^..\n"

func Test_parseInput(t *testing.T) {
	reader := strings.NewReader(testData)
	inp, err := parseInput(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if inp.startX != 2 || inp.startY != 1 {
		t.Errorf("unexpected start position: got (%d,%d), want (2,1)", inp.startX, inp.startY)
	}
	if inp.startNode == nil {
		t.Errorf("unexpected start node: got nil, want non-nil")
	}
	if inp.width != 5 {
		t.Errorf("unexpected width: got %d, want 5", inp.width)
	}
	if inp.height != 4 {
		t.Errorf("unexpected height: got %d, want 4", inp.height)
	}
}

func Benchmark_parseInput(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {
		for b.Loop() {
			reader := strings.NewReader(testData)
			inp, err := parseInput(reader)
			if err != nil {
				b.Fatalf("Unexpected error: %v", err)
			}
			if inp.startNode == nil || inp.startX != 2 || inp.startY != 1 || inp.width != 5 || inp.height != 4 {
				b.Fatalf("unexpected parsed data: start(%d,%d) width %d height %d", inp.startX, inp.startY, inp.width, inp.height)
			}
		}
	})
}
