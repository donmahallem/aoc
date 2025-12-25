package day13

import (
	"io"
)

func validateAxis(items []int, center int) bool {
	for lower, upper := center-1, center; lower >= 0 && upper < len(items); lower, upper = lower-1, upper+1 {
		if items[lower] != items[upper] {
			return false
		}
	}
	return true
}

func findAxis(axisData []int) (int, bool) {
	for centerIdx := 1; centerIdx < len(axisData); centerIdx++ {
		if axisData[centerIdx-1] == axisData[centerIdx] {
			if validateAxis(axisData, centerIdx) {
				return centerIdx, true
			}
		}
	}
	return -1, false
}

func Part1(in io.Reader) (int, error) {
	start, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	accum := 0
	for _, block := range start {
		if rowAxis, ok := findAxis(block.Rows); ok {
			accum += rowAxis * 100
			continue
		}
		if colAxis, ok := findAxis(block.Cols); ok {
			accum += colAxis
			continue
		}
	}
	return accum, nil
}
