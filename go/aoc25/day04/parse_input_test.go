package day04

import (
	"strings"
	"testing"
)

var testData = `@.@
..@
.@.
`

func Test_parseInput(t *testing.T) {
	reader := strings.NewReader(testData)
	got, width, height := parseInput(reader)

	if len(got) != 9 {
		t.Errorf("Expected len of %d, got %d", 9, len(got))
	}

	if width != 3 || height != 3 {
		t.Errorf("Expected width and height to be 10,10 got %d,%d", width, height)
	}

}
