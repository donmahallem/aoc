package day02_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day02"
)

const testData string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

var testDataSlices []string = []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}

func TestParseLine(t *testing.T) {
	t.Run("test line 0", func(t *testing.T) {
		expected := []day02.Block{{4, 0, 3}, {1, 2, 6}, {0, 2, 0}}
		data := []byte(testDataSlices[0])
		game, blocks := day02.ParseLine(&data)
		if game != 1 || !slices.Equal(expected, blocks) {
			t.Errorf(`Expected %d - %v to match %d - %v`, game, blocks, 1, expected)
		}
	})
	t.Run("test line 1", func(t *testing.T) {
		expected := []day02.Block{{0, 2, 1}, {1, 3, 4}, {0, 1, 1}}
		data := []byte(testDataSlices[1])
		game, blocks := day02.ParseLine(&data)
		if game != 2 || !slices.Equal(expected, blocks) {
			t.Errorf(`Expected %d - %v to match %d - %v`, game, blocks, 2, expected)
		}
	})
}

func TestValidateBlocks(t *testing.T) {
	validLines := [4]int{0, 1, 4}
	invalidLines := [2]int{2, 3}
	for _, idx := range validLines {
		t.Run(fmt.Sprintf("Test line %d", idx), func(t *testing.T) {
			data := []byte(testDataSlices[idx])
			_, blocks := day02.ParseLine(&data)
			if !day02.ValidateBlocks(&blocks) {
				t.Errorf(`Expected to be valid block`)
			}
		})
	}
	for _, idx := range invalidLines {
		t.Run(fmt.Sprintf("Test line %d", idx), func(t *testing.T) {
			data := []byte(testDataSlices[idx])
			_, blocks := day02.ParseLine(&data)
			if day02.ValidateBlocks(&blocks) {
				t.Errorf(`Expected to be invalid block`)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	result := day02.Part1(strings.NewReader(testData))
	if result != 8 {
		t.Errorf(`Expected %d to be %d`, result, 8)
	}
}

func BenchmarkPart1(b *testing.B) {
	for b.Loop() {
		day02.Part1(strings.NewReader(testData))
	}
}
