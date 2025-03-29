package day03

import (
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

func Part2(in io.Reader) int {
	field, _ := aoc_utils.LoadField(in)
	parts, matches := FindObjects(field)
	pairs := PairObjects(parts, matches)
	summe := 0
	for pairIdx := range len(pairs) {
		if pairs[pairIdx].Part.PartType == '*' && len(pairs[pairIdx].Matches) == 2 {
			summe += pairs[pairIdx].Matches[0].Value * pairs[pairIdx].Matches[1].Value
		}
	}
	return summe
}
