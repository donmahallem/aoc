package day12

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed sample.txt
var testData string

func Test_parseInput(t *testing.T) {
	t.Run("testData", func(t *testing.T) {
		reader := strings.NewReader(testData)
		lines, err := parseInput(reader, 1)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		const expectedLen int = 6
		if len(lines) != expectedLen {
			t.Errorf(`Expected %d lines to be parsed, got %d`, expectedLen, len(lines))
		}
	})
}
