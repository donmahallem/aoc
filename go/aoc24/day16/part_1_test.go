package day16

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"
)

//go:embed sample1.txt
var testData1 string

//go:embed sample2.txt
var testData2 string

func TestParseInput(t *testing.T) {
	inputData, err := parseInput(strings.NewReader(testData1))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(inputData.Field) != 15 {
		t.Errorf(`Expected %d to match %d`, len(inputData.Field), 15)
	}
	if inputData.Start.Y != 13 || inputData.Start.X != 1 {
		t.Errorf(`Expected %v to match [4,4]`, inputData.Start)
	}
	if inputData.End.Y != 1 || inputData.End.X != 13 {
		t.Errorf(`Expected %v to match [4,4]`, inputData.End)
	}
}

func fieldToCsv(f *field) {
	p := point{}
	for y := range len(*f) {
		p.Y = int16(y)
		for x := range len((*f)[y]) {
			p.X = int16(x)
			if x > 0 {
				fmt.Print(",")
			}
			fmt.Printf("%d", (*f)[y][x])
		}
		fmt.Println()
	}
}

func TestFindShortestPath(t *testing.T) {
	inputData, err := parseInput(strings.NewReader(testData1))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	calculatePathValues(&inputData.Field, &inputData.Start)
	if inputData.Field[inputData.End.Y][inputData.End.X] != 7036 {
		t.Errorf(`Expected %d to match 7036`, inputData.Field[inputData.End.Y][inputData.End.X])
	}
}

func TestCalculatePathValues_Invalid(t *testing.T) {
	// empty field should not panic
	f := field{}
	p := point{X: 0, Y: 0}
	calculatePathValues(&f, &p)
	// start outside of field should not panic
	f2 := make(field, 1)
	f2[0] = make([]int, 1)
	p2 := point{X: 10, Y: 10}
	calculatePathValues(&f2, &p2)
}
