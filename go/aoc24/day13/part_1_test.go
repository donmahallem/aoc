package day13_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day13"
)

const testData string = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

func TestOutOfBoundsShouldBeInside(t *testing.T) {
	test := []byte(" X+94, Y+34")
	expected := [2]float64{94, 34}
	if res := day13.ParseButton(test); res != expected {
		t.Errorf(`Expected %v to match %v`, res, expected)
	}
}

func TestFindNeighbours(t *testing.T) {
	test := day13.LoadFile(strings.NewReader(testData))
	if len(test) != 4 {
		t.Errorf(`Expected %d to match %d`, len(test), 4)
	}
}
