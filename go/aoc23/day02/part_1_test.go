package day02_test

import (
	"slices"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day02"
)

var testData []string = []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}

func TestParseLine(t *testing.T) {
	expected := []day02.Block{{4, 0, 3}, {1, 2, 6}, {0, 2, 0}}
	data := []byte(testData[0])
	game, blocks := day02.ParseLine(&data)
	if game != 1 || !slices.Equal(expected, blocks) {
		t.Errorf(`Expected %d - %v to match %d - %v`, game, blocks, 1, expected)
	}
	expected = []day02.Block{{0, 2, 1}, {1, 3, 4}, {0, 1, 1}}
	data = []byte(testData[1])
	game, blocks = day02.ParseLine(&data)
	if game != 2 || !slices.Equal(expected, blocks) {
		t.Errorf(`Expected %d - %v to match %d - %v`, game, blocks, 2, expected)
	}
}

func TestValidateBlocks(t *testing.T) {
	data := []byte(testData[0])
	_, blocks := day02.ParseLine(&data)
	if !day02.ValidateBlocks(&blocks) {
		t.Errorf(`Expected to be valid block`)
	}
	data = []byte(testData[1])
	_, blocks = day02.ParseLine(&data)
	if !day02.ValidateBlocks(&blocks) {
		t.Errorf(`Expected to be invalid block`)
	}
	data = []byte(testData[2])
	_, blocks = day02.ParseLine(&data)
	if day02.ValidateBlocks(&blocks) {
		t.Errorf(`Expected to be valid block`)
	}
	data = []byte(testData[3])
	_, blocks = day02.ParseLine(&data)
	if day02.ValidateBlocks(&blocks) {
		t.Errorf(`Expected to be valid block`)
	}
	data = []byte(testData[4])
	_, blocks = day02.ParseLine(&data)
	if !day02.ValidateBlocks(&blocks) {
		t.Errorf(`Expected to be valid block`)
	}
}
