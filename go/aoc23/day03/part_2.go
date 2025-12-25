package day03

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func Part2(in io.Reader) (int, error) {
	field, err := aoc_utils.LoadField[int16, byte](in)
	if err != nil {
		return 0, err
	}
	parts, matches := findObjects(*field)
	pairs := pairObjects(parts, matches)
	summe := 0
	for pairIdx := range len(pairs) {
		if pairs[pairIdx].Part.PartType == '*' && len(pairs[pairIdx].Matches) == 2 {
			summe += pairs[pairIdx].Matches[0].Value * pairs[pairIdx].Matches[1].Value
		}
	}
	return summe, nil
}
