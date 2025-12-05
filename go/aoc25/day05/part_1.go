package day05

import (
	"io"
)

func findRange(ranges []validRange, value uint64) bool {
	// divide and conquer search as the ranges are non-overlapping and sorted
	low := 0
	high := len(ranges) - 1
	for low <= high {
		mid := (low + high) / 2
		if ranges[mid].Contains(value) {
			return true
		} else if value < ranges[mid].Min {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func Part1(in io.Reader) int {
	inp := parseInput(in)
	compressValidRanges(&inp.validRanges)
	validIngredientCount := 0
	for ingredient := range inp.ingredients {
		if findRange(inp.validRanges, inp.ingredients[ingredient]) {
			validIngredientCount++
		}
	}
	return validIngredientCount
}
