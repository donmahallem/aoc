package day05

import (
	"io"
)

func Part2(in io.Reader) (int, error) {
	inp, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	compressValidRanges(&inp.validRanges)
	validIngredientCount := 0
	for interval := range inp.validRanges {
		validIngredientCount += int(inp.validRanges[interval].Max - inp.validRanges[interval].Min + 1)
	}
	return validIngredientCount, nil
}
