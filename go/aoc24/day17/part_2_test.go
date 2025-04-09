package day17_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day17"
)

const testData2 string = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

func TestSearchForNeutralElement(t *testing.T) {
	reg := day17.Register{0, 0, 0}
	prog_target := day17.Program{1}
	prog := day17.Program{2, 1}
	for !day17.SearchForNeutralElement(&reg, &prog, &prog_target) {
		reg[0]++
	}
	if reg[0] > 0 {
		t.Errorf(`Expected %d to match false`, reg[0])
	}
	fmt.Printf("%v - %v - %v\n", reg, prog_target, prog)
}

func TestPart2_testData2(t *testing.T) {
	result := day17.Part2(strings.NewReader(testData2))
	if result != 117440 {
		t.Errorf(`Expected %d to match 117440`, result)
	}
}
