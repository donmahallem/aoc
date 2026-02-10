package day15

import (
	"strings"
	"testing"
)

func Test_parseInput(t *testing.T) {
	data, err := parseInput(strings.NewReader(testDataBig), false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(data.Field) != 10 {
		t.Errorf(`Expected %d to match %d`, len(data.Field), 10)
	}
	if data.Player.Y != 4 || data.Player.X != 4 {
		t.Errorf(`Expected %v to match [4,4]`, data.Player)
	}
	if len(data.Movements) != 700 {
		t.Errorf(`Expected %d to match %d`, len(data.Movements), 700)
	}
}
func Test_parseInputDoubleWide(t *testing.T) {
	data, err := parseInput(strings.NewReader(testDataBig), true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(data.Field) != 10 || len(data.Field[0]) != 20 {
		t.Errorf(`Expected field to have size(10,20) not (%d,%d)`, len(data.Field), len(data.Field[0]))
	}
	if data.Player.Y != 4 || data.Player.X != 8 {
		t.Errorf(`Expected %v to match [4,8]`, data.Player)
	}
	if len(data.Movements) != 700 {
		t.Errorf(`Expected %d to match %d`, len(data.Movements), 700)
	}
}
