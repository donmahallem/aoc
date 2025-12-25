package day11

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed sample.txt
var testData string

func Test_parseInput(t *testing.T) {
	t.Run("offset 1", func(t *testing.T) {
		reader := strings.NewReader(testData)
		galaxies, err := parseInput(reader, 1)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		const expectedLen int = 9
		if len(galaxies) != expectedLen {
			t.Errorf(`Expected %d galaxies to be parsed, got %d`, expectedLen, len(galaxies))
		}
	})
	t.Run("offset 100", func(t *testing.T) {
		reader := strings.NewReader(testData)
		galaxies, err := parseInput(reader, 100)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		const expectedLen int = 9
		if len(galaxies) != expectedLen {
			t.Errorf(`Expected %d galaxies to be parsed, got %d`, expectedLen, len(galaxies))
		}
	})
}
