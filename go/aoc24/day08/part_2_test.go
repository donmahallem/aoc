package day08

import (
	"io"
	"strings"
	"testing"
)

func Test24Day08Part2(t *testing.T) {

	result, err := Part2(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 34 {
		t.Errorf(`Expected %d to match %d`, result, 34)
	}
}

func BenchmarkPart2(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		Part2(data)
	}
}
