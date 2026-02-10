package day03

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func Part1(in io.Reader) (int, error) {
	field, err := aoc_utils.LoadField[int16, byte](in)
	if err != nil {
		return 0, err
	}
	if field == nil {
		return 0, aoc_utils.NewParseError("empty field", nil)
	}
	parts, matches := findObjects(*field)
	pairs := pairObjects(parts, matches)
	summe := 0
	for pairIdx := range len(pairs) {
		for matchIdx := range len(pairs[pairIdx].Matches) {
			summe += pairs[pairIdx].Matches[matchIdx].Value
		}
	}
	return summe, nil
}
