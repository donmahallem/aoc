package day03

import (
	_ "embed"
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

//go:embed sample.txt
var testData string

func TestFindParts(t *testing.T) {
	result, _ := aoc_utils.LoadField[int16, byte](strings.NewReader(testData))
	parts, matches := findObjects(*result)
	if len(parts) != 6 {
		t.Errorf(`Expected to be 6 parts`)
	}
	expectedParts := []part{*newPart('*', 3, 1),
		*newPart('#', 6, 3),
		*newPart('*', 3, 4),
		*newPart('+', 5, 5),
		*newPart('$', 3, 8),
		*newPart('*', 5, 8)}
	for _, exp := range expectedParts {
		if !slices.Contains(parts, exp) {
			t.Errorf(`Expected to contain %v. Contains %v`, exp, parts)
		}
	}
	expectedMatches := []number{
		*newNumber(467, 0, 2, 0),
		*newNumber(114, 5, 7, 0),
		*newNumber(35, 2, 3, 2),
		*newNumber(633, 6, 8, 2),
		*newNumber(617, 0, 2, 4),
		*newNumber(58, 7, 8, 5),
		*newNumber(592, 2, 4, 6),
		*newNumber(755, 6, 8, 7),
		*newNumber(664, 1, 3, 9),
		*newNumber(598, 5, 7, 9),
		*newNumber(123, 7, 9, 4)}
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
	result, _ := aoc_utils.LoadField[int16, byte](strings.NewReader(testData))
	parts, matches := findObjects(*result)
	pairs := pairObjects(parts, matches)
	if len(pairs) != len(parts) {
		t.Errorf(`Expected %d to be %v`, len(parts), pairs)
	}
}
