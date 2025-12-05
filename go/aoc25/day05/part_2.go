package day05

import (
	"io"
)

func Part2(in io.Reader) int {
	inp := parseInput(in)
	compressValidRanges(&inp.validRanges)
	validIngredientCount := 0
	for interval := range inp.validRanges {
		validIngredientCount += int(inp.validRanges[interval].Max - inp.validRanges[interval].Min + 1)
	}
	return validIngredientCount
}
