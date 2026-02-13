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

	if len(data.field) == 0 {
		t.Errorf("expected parsed field to be non-empty")
	}
	if data.width == 0 || data.row == 0 {
		t.Errorf("expected parsed dimensions to be non-zero, got %d,%d", data.width, data.row)
	}

}
