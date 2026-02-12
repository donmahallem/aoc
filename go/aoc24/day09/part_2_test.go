package day09

import (
	"slices"
	"testing"
)

func TestFindEmptySpace(t *testing.T) {
	test := []int16{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}
	var lookFor int = 1
	var endIdx int = 10
	if start, end, ok := FindEmptySpace(&test, &lookFor, &endIdx); start != 1 && end != 3 && !ok {
		t.Errorf(`Expected (%d,%d,%t) and got (%d,%d,%t)`, 1, 3, true, start, end, ok)
	}
	lookFor = 2
	endIdx = 10
	if start, end, ok := FindEmptySpace(&test, &lookFor, &endIdx); start != 2 && end != 4 && !ok {
		t.Errorf(`Expected (%d,%d,%t) and got (%d,%d,%t)`, 1, 4, true, start, end, ok)
	}
	lookFor = 6
	endIdx = 10
	if start, end, ok := FindEmptySpace(&test, &lookFor, &endIdx); start != -1 && end != -1 && ok {
		t.Errorf(`Expected (%d,%d,%t) and got (%d,%d,%t)`, -1, -1, false, start, end, ok)
	}
}

func TestFindBlockWithNegStart(t *testing.T) {
	test := []int16{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}
	maxIdx := -1
	var id int16 = 2
	if start, end, ok := FindBlock(&test, &id, &maxIdx); start != 10 && end != 15 && !ok {
		t.Errorf(`Expected (%d,%d,%t) and got (%d,%d,%t)`, 10, 15, true, start, end, ok)
	}
	id = 1
	if start, end, ok := FindBlock(&test, &id, &maxIdx); start != 3 && end != 6 && !ok {
		t.Errorf(`Expected (%d,%d,%t) and got (%d,%d,%t)`, 3, 6, true, start, end, ok)
	}
	id = 6
	if start, end, ok := FindBlock(&test, &id, &maxIdx); start != -1 && end != -1 && ok {
		t.Errorf(`Expected (%d,%d,%t) and got (%d,%d,%t)`, -1, -1, true, start, end, ok)
	}
}
func TestCompactLess(t *testing.T) {
	test := []byte("2333133121414131402")
	expandedData, err := convertInput(test)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	CompactLess(&expandedData)
	expected := []int16{0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, -1, 4, 4, -1, 3, 3, 3, -1, -1, -1, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, -1, -1, -1, -1, 8, 8, 8, 8, -1, -1}
	if !slices.Equal(expandedData, expected) {
		t.Errorf(`Expected %v to be %v`, expandedData, expected)
	}
	if checkSum := checkSum(&expandedData); checkSum != 2858 {
		t.Errorf(`Expected checksum %d to be %d`, checkSum, 2858)
	}
}
