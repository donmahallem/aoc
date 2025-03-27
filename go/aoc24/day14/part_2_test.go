package day14_test

import (
	"testing"

	"github.com/donmahallem/aoc/aoc24/day14"
)

func TestGetPointDeviation_2(t *testing.T) {
	tests := []day14.Point[int]{*day14.NewPoint(3, 0),
		*day14.NewPoint(6, 2),
		*day14.NewPoint(6, 6)}
	var expected float64 = 3.0 / 4.0
	if res := day14.GetPointDeviation(&tests); res != expected {
		t.Errorf(`Expected %f to match %f`, res, expected)
	}
}
func TestGetDiff(t *testing.T) {
	tests := []day14.Robot{*day14.NewRobot(5, 1, 3, -3),
		*day14.NewRobot(9, 1, 3, -3),
		*day14.NewRobot(12, 1, 3, -3),
		*day14.NewRobot(5, 2, 3, -3),
		*day14.NewRobot(8, 4, 3, -3),
		*day14.NewRobot(9, 4, 3, -3)}
	var expected float64 = 9.0 / 4.0
	if res := day14.GetDiff(&tests); res != expected {
		t.Errorf(`Expected %f to match %f`, res, expected)
	}
}
func TestGetDiff_Other(t *testing.T) {
	tests := []day14.Robot{*day14.NewRobot(3, 0, 3, -3),
		*day14.NewRobot(6, 2, 3, -3),
		*day14.NewRobot(7, 2, 3, -3),
		*day14.NewRobot(3, 6, 3, -3),
		*day14.NewRobot(7, 6, 3, -3),
		*day14.NewRobot(10, 6, 3, -3)}
	if res := day14.GetDiff(&tests); res != -1.0/6.0 {
		t.Errorf(`Expected %f to match %f`, res, -0.166667)
	}
}
func TestFindLowestDeviation(t *testing.T) {
	tests := []day14.Robot{*day14.NewRobot(5, 1, 3, -3),
		*day14.NewRobot(9, 1, 3, -3),
		*day14.NewRobot(12, 1, 3, -3),
		*day14.NewRobot(5, 2, 3, -3),
		*day14.NewRobot(8, 4, 3, -3),
		*day14.NewRobot(9, 4, 3, -3)}
	if res := day14.FindLowestDeviation(tests, 10, 11, 7); res != 5135 {
		t.Errorf(`Expected %d to match %d`, res, 4)
	}
}

func TestStep(t *testing.T) {
	tests := []day14.Robot{*day14.NewRobot(5, 1, 3, -3)}
	expected := *day14.NewRobot(8, 5, 3, -3)
	width, height := 11, 7
	day14.Step(&tests, &width, &height)
	if tests[0] != expected {
		t.Errorf(`Expected %v to match %v`, tests[0], expected)
	}
}
