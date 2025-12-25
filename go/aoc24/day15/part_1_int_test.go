package day15

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed sample_1.txt
var testDataBig string

//go:embed sample_2.txt
var testDataSmall string

func TestFindNextEmptyCellOffset_SingleBlock(t *testing.T) {
	data, err := parseInput(strings.NewReader(testDataBig), false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	testPlayer := player{Y: 4, X: 4}
	testMove := move{Y: 0, X: -1}
	off, ok := findNextEmptyCellOffset(&data.Field, &testPlayer, &testMove)
	if !ok || off != 1 {
		t.Errorf(`Expected offset to be ok and 1 not %d`, off)
	}
}
func TestFindNextEmptyCellOffset_DoubleBlock(t *testing.T) {
	data, err := parseInput(strings.NewReader(testDataBig), false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	testPlayer := player{Y: 3, X: 4}
	testMove := move{Y: 0, X: -1}
	off, ok := findNextEmptyCellOffset(&data.Field, &testPlayer, &testMove)
	if !ok || off != 2 {
		t.Errorf(`Expected offset to be ok and 2 not %d`, off)
	}
}
func TestFindNextEmptyCellOffset_TripleBlock(t *testing.T) {
	data, err := parseInput(strings.NewReader(testDataBig), false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	testPlayer := player{Y: 3, X: 4}
	data.Field[4][4] = CELL_BOX
	data.Field[5][4] = CELL_BOX
	testMove := move{Y: 1, X: 0}
	off, ok := findNextEmptyCellOffset(&data.Field, &testPlayer, &testMove)
	if !ok || off != 3 {
		t.Errorf(`Expected offset to be ok and 3 not %d`, off)
	}
}
func TestFindNextEmptyCellOffset_QuadBlock(t *testing.T) {
	// Testing 4 Blocks Space Wall
	data, err := parseInput(strings.NewReader(testDataBig), false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	testPlayer := player{Y: 3, X: 4}
	data.Field[4][4] = CELL_BOX
	data.Field[5][4] = CELL_BOX
	data.Field[6][4] = CELL_BOX
	data.Field[7][4] = CELL_BOX
	data.Field[8][4] = CELL_EMPTY
	testMove := move{Y: 1, X: 0}
	off, ok := findNextEmptyCellOffset(&data.Field, &testPlayer, &testMove)
	if !ok || off != 4 {
		t.Errorf(`Expected offset to be ok and 4 not %d`, off)
	}
}
func TestFindNextEmptyCellOffset_BlocksAgainstWall(t *testing.T) {
	// Testing 5 Blocks Wall
	data, err := parseInput(strings.NewReader(testDataBig), false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	testPlayer := player{Y: 3, X: 4}
	data.Field[4][4] = CELL_BOX
	data.Field[5][4] = CELL_BOX
	data.Field[6][4] = CELL_BOX
	data.Field[7][4] = CELL_BOX
	data.Field[8][4] = CELL_BOX
	testMove := move{Y: 1, X: 0}
	off, ok := findNextEmptyCellOffset(&data.Field, &testPlayer, &testMove)
	if ok {
		t.Errorf(`Expected offset to be ok and 4 not %d`, off)
	}
}
