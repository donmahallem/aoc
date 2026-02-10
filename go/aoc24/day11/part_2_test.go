package day11

import (
	"io"
	"strings"
	"testing"
)

func TestPart2(t *testing.T) {
	result, err := Part2(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 65601038650482 {
		t.Errorf(`Expected %d to contain %d`, result, 65601038650482)
	}
}

func BenchmarkPart2(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		Part2(data)
	}
}
