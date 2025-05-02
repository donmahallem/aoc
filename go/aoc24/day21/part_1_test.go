package day21_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day21"
)

const testData string = `029A
980A
179A
456A
379A`

func TestIterateInput(t *testing.T) {
	test := day21.IterateInput(strings.NewReader(testData))
	items := make([][]byte, 0, 5)
	for item := range test {
		items = append(items, item)
	}
	if len(items) != 5 {
		t.Errorf(`Expected %v to match 5`, items)
	}
}

func TestParseIntValue(t *testing.T) {
	test := day21.IterateInput(strings.NewReader(testData))
	items := make([]uint, 0, 5)
	for item := range test {
		items = append(items, day21.ParseIntValue(&item))
	}
	if len(items) != 5 {
		t.Errorf(`Expected %v to match 5`, items)
	}
	if items[0] != 29 {
		t.Errorf(`Expected %d to match 29`, items[0])
	}
	if items[1] != 980 {
		t.Errorf(`Expected %d to match 980`, items[1])
	}
	if items[2] != 179 {
		t.Errorf(`Expected %d to match 179`, items[2])
	}
	if items[3] != 456 {
		t.Errorf(`Expected %d to match 456`, items[3])
	}
	if items[4] != 379 {
		t.Errorf(`Expected %d to match 379`, items[4])
	}
}
func TestCalculateMoves(t *testing.T) {
	t.Run("Test depth 3", func(t *testing.T) {
		test := day21.CalculateMoves(strings.NewReader(testData), 3)
		if test != 126384 {
			t.Errorf(`Expected %d to match 126384`, test)
		}
	})
	t.Run("Test depth 4", func(t *testing.T) {
		test := day21.CalculateMoves(strings.NewReader(testData), 4)
		if test != 310188 {
			t.Errorf(`Expected %d to match 310188`, test)
		}
	})
	t.Run("Test depth 5", func(t *testing.T) {
		test := day21.CalculateMoves(strings.NewReader(testData), 5)
		if test != 757754 {
			t.Errorf(`Expected %d to match 757754`, test)
		}
	})
}

func TestPart1(t *testing.T) {
	test := day21.Part1(strings.NewReader(testData))
	if test != 126384 {
		t.Errorf(`Expected %d to match 126384`, test)
	}
}
