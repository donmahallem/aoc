package day21_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day21"
)

func TestPart2(t *testing.T) {
	test, err := day21.Part2(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if test != 154115708116294 {
		t.Errorf(`Expected %d to match 154115708116294`, test)
	}
}

func BenchmarkPart2(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day21.Part2(data)
	}
}
