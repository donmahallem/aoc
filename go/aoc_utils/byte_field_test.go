package aoc_utils

import (
	"strings"
	"testing"
)

func TestLoadField_UnevenLines(t *testing.T) {
	// first line length 3, second line length 2 -> should return parse error
	input := "abc\nxy\n"
	_, err := LoadFieldWithOffset[int, byte](strings.NewReader(input), 0)
	if err == nil {
		t.Fatalf("expected error for uneven lines, got nil")
	}
}
