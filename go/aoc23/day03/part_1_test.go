package day03_test

import (
	"io"
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day03"
	"github.com/donmahallem/aoc/aoc_utils"
)

const testData = `467..114..
...*......
..35..633.
......#...
617*...123
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestFindParts(t *testing.T) {
	result, _ := aoc_utils.LoadField[int16](strings.NewReader(testData))
	parts, matches := day03.FindObjects(*result)
	if len(parts) != 6 {
		t.Errorf(`Expected to be 6 parts`)
	}
	expectedParts := []day03.Part{*day03.NewPart('*', 3, 1),
		*day03.NewPart('#', 6, 3),
		*day03.NewPart('*', 3, 4),
		*day03.NewPart('+', 5, 5),
		*day03.NewPart('$', 3, 8),
		*day03.NewPart('*', 5, 8)}
	for _, exp := range expectedParts {
		if !slices.Contains(parts, exp) {
			t.Errorf(`Expected to contain %v. Contains %v`, exp, parts)
		}
	}
	expectedMatches := []day03.Number{
		*day03.NewNumber(467, 0, 2, 0),
		*day03.NewNumber(114, 5, 7, 0),
		*day03.NewNumber(35, 2, 3, 2),
		*day03.NewNumber(633, 6, 8, 2),
		*day03.NewNumber(617, 0, 2, 4),
		*day03.NewNumber(58, 7, 8, 5),
		*day03.NewNumber(592, 2, 4, 6),
		*day03.NewNumber(755, 6, 8, 7),
		*day03.NewNumber(664, 1, 3, 9),
		*day03.NewNumber(598, 5, 7, 9),
		*day03.NewNumber(123, 7, 9, 4)}
	if len(matches) != 11 {
		t.Errorf(`Expected to be 11 parts. Got %v`, matches)
	}
	for _, exp := range expectedMatches {
		if !slices.Contains(matches, exp) {
			t.Errorf(`Expected to contain %v. Contains %v`, exp, matches)
		}
	}
}

func TestPairObjects(t *testing.T) {
	result, _ := aoc_utils.LoadField[int16](strings.NewReader(testData))
	parts, matches := day03.FindObjects(*result)
	pairs := day03.PairObjects(parts, matches)
	if len(pairs) != len(parts) {
		t.Errorf(`Expected %d to be %v`, result, pairs)
	}
}

func TestPart1(t *testing.T) {
	t.Run("test sample data", func(t *testing.T) {
		result := day03.Part1(strings.NewReader(testData))
		if result != 4484 {
			t.Errorf(`Expected %d to be %d`, result, 4484)
		}
	})
	t.Run("test sample data", func(t *testing.T) {
		result := day03.Part1(strings.NewReader("....@123\n456....."))
		if result != 123 {
			t.Errorf(`Expected %d to be %d`, result, 123)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day03.Part1(reader)
	}
}
