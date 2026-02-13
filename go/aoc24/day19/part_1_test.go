package day19

import (
	"testing"
)

const testPatterns string = "r, wr, b, g, bwu, rb, gb, br"

func TestParseFirstLine(t *testing.T) {
	testInput := testPatterns
	points, keyLen, err := ParseFirstLine(testInput)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(points) != 8 {
		t.Errorf(`Expected %d to match 8`, len(points))
	}
	if keyLen != 3 {
		t.Errorf(`Expected %d to match 3`, keyLen)
	}
}
