package day17_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day17"
)

const testData1 string = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

func TestParseInput(t *testing.T) {
	registers, program := day17.ParseInput(strings.NewReader(testData1))
	expectedRegisters := day17.Register{729, 0, 0}
	expectedProgram := day17.Program{0, 1, 5, 4, 3, 0}
	if registers != expectedRegisters {
		t.Errorf(`Expected %v to match %v`, registers, expectedRegisters)
	}
	if !slices.Equal(program, expectedProgram) {
		t.Errorf(`Expected %v to match %v`, program, expectedProgram)
	}
}

func TestPart1_testData1(t *testing.T) {
	result := day17.Part1(strings.NewReader(testData1))
	expected := []int{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}
	if !slices.Equal(result, expected) {
		t.Errorf(`Expected %v to match %v`, result, expected)
	}
}
