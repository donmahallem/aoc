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
	data, err := parseInput(reader)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(data.field) != 9 {
		t.Errorf("Expected len of %d, got %d", 9, len(data.field))
	}

	if data.width != 3 || data.row != 3 {
		t.Errorf("Expected width and height to be 3,3 got %d,%d", data.width, data.row)
	}

}
