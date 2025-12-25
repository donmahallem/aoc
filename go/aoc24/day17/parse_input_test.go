package day17

import (
	"slices"
	"strings"
	"testing"

	_ "embed"
)

//go:embed sample1.txt
var testData1 string

func TestParseInput(t *testing.T) {
	data, err := parseInput(strings.NewReader(testData1))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedRegisters := Register{729, 0, 0}
	expectedProgram := Program{0, 1, 5, 4, 3, 0}
	if data.Register != expectedRegisters {
		t.Errorf(`Expected %v to match %v`, data.Register, expectedRegisters)
	}
	if !slices.Equal(data.Program, expectedProgram) {
		t.Errorf(`Expected %v to match %v`, data.Program, expectedProgram)
	}
}
