package day08

import (
	_ "embed"
	"io"
	"strings"
	"testing"
)

//go:embed sample.txt
var testData string

func Test24Day08Part1(t *testing.T) {

	result, err := Part1(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 14 {
		t.Errorf(`Expected %d to match %d`, result, 14)
	}
}

func BenchmarkPart1(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		Part1(data)
	}
}
