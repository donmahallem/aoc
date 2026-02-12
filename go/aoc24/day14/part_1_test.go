package day14_test

import (
	_ "embed"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day14"
)

func TestParseLine(t *testing.T) {
	testString := []byte("p=0,4 v=3,-3")
	r, err := day14.ParseLine(testString)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	test := *r
	expected := *day14.NewRobot(0, 4, 3, -3)
	if expected != test {
		t.Errorf(`Expected %v to match %v`, test, expected)
	}
}

func TestParseLineLongValue(t *testing.T) {
	testString := []byte("p=-20,145 v=3,-253")
	r, err := day14.ParseLine(testString)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	test := *r
	expected := *day14.NewRobot(-20, 145, 3, -253)
	if expected != test {
		t.Errorf(`Expected %v to match %v`, test, expected)
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
