package day14_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day14"
)

const testData string = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func TestParseLine(t *testing.T) {
	testString := []byte("p=0,4 v=3,-3")
	test := day14.ParseLine(&testString)
	expected := *day14.NewRobot(0, 4, 3, -3)
	if expected != test {
		t.Errorf(`Expected %v to match %v`, test, expected)
	}
}

func TestParseLineLongValue(t *testing.T) {
	testString := []byte("p=-20,145 v=3,-253")
	test := day14.ParseLine(&testString)
	expected := *day14.NewRobot(-20, 145, 3, -253)
	if expected != test {
		t.Errorf(`Expected %v to match %v`, test, expected)
	}
}
func TestLoadFile(t *testing.T) {
	test := day14.LoadFile(strings.NewReader(testData))
	expected := []*day14.Robot{day14.NewRobot(0, 4, 3, -3),
		day14.NewRobot(6, 3, -1, -3),
		day14.NewRobot(10, 3, -1, 2),
		day14.NewRobot(2, 0, 2, -1),
		day14.NewRobot(0, 0, 1, 3),
		day14.NewRobot(3, 0, -2, -2),
		day14.NewRobot(7, 6, -1, -3),
		day14.NewRobot(3, 0, -1, -2),
		day14.NewRobot(9, 3, 2, 3),
		day14.NewRobot(7, 3, -1, 2),
		day14.NewRobot(2, 4, 2, -3),
		day14.NewRobot(9, 5, -3, -3)}
	for _, exp := range expected {
		if !slices.Contains(test, *exp) {
			t.Errorf(`Expected %v to match %d`, *exp, 12)
		}
	}
}

func TestCountQuadrant(t *testing.T) {
	test := day14.LoadFile(strings.NewReader(testData))
	testSum := day14.CountQuadrant(&test, 100, 11, 7)
	if testSum != 12 {
		t.Errorf(`Expected %v to match %v`, testSum, 12)
	}
}

func TestCalculateQuadrant(t *testing.T) {
	test := *day14.NewRobot(0, 1, 3, -3)
	testSum := day14.CalculateQuadrant(&test, 0, 5, 5)
	if testSum != 0 {
		t.Errorf(`Expected %v to match %v`, testSum, 0)
	}
	test = *day14.NewRobot(2, 1, 3, -3)
	testSum = day14.CalculateQuadrant(&test, 0, 5, 5)
	if testSum != -1 {
		t.Errorf(`Expected %v to match %v`, testSum, -1)
	}
	test = *day14.NewRobot(4, 4, 3, -3)
	testSum = day14.CalculateQuadrant(&test, 0, 5, 5)
	if testSum != 3 {
		t.Errorf(`Expected %v to match %v`, testSum, 3)
	}
	test = *day14.NewRobot(2, 4, 2, -3)
	testSum = day14.CalculateQuadrant(&test, 5, 11, 7)
	if testSum != -1 {
		t.Errorf(`Expected %v to match %v`, testSum, -1)
	}
}

func BenchmarkPart1(b *testing.B) {
	for b.Loop() {
		day14.Part1(strings.NewReader(testData))
	}
}
