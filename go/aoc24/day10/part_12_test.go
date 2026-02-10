package day10

import (
	"strings"
	"testing"
)

const testData string = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestWalkDepth(t *testing.T) {
	f1, err := loadField(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	result := make(map[uint16]struct{})
	walkDepth(*f1, 2, 0, result)
	if len(result) != 5 {
		t.Errorf(`Expected %v to not match %d`, len(result), 5)
	}
	result = make(map[uint16]struct{})
	walkDepth(*f1, 4, 0, result)
	if len(result) != 6 {
		t.Errorf(`Expected %v to not match %d`, len(result), 5)
	}
}
func TestSearchAll(t *testing.T) {
	f1, err := loadField(strings.NewReader(testData))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	result := searchAll(*f1)
	if result != 36 {
		t.Errorf(`Expected %d to not match %d`, result, 36)
	}
}
