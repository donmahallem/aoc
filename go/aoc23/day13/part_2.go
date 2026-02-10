package day13

import (
	"io"
)

func oneBitApart(a, b int) bool {
	diff := a ^ b
	return (diff != 0 && (diff&(diff-1)) == 0)
}

// checks whether the axisData is symmetric around center allowing for one bit correction
func validateAxis2(axisData []int, center int) bool {
	corrected := 0
	for lower, upper := center-1, center; lower >= 0 && upper < len(axisData); lower, upper = lower-1, upper+1 {
		if axisData[lower] != axisData[upper] {
			if corrected == 0 {
				if oneBitApart(axisData[lower], axisData[upper]) {
					corrected++
					continue
				}
			}
			return false
		}
	}
	return corrected == 1
}
func findAxis2(axisData []int) (int, bool) {
	for centerIdx := 1; centerIdx < len(axisData); centerIdx++ {
		if validateAxis2(axisData, centerIdx) {
			return centerIdx, true
		}
	}
	return -1, false
}

func Part2(in io.Reader) (int, error) {
	start, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	accum := 0
	for _, block := range start {
		if rowAxis, ok := findAxis2(block.Rows); ok {
			accum += rowAxis * 100
			continue
		}
		if colAxis, ok := findAxis2(block.Cols); ok {
			accum += colAxis
			continue
		}
	}
	return accum, nil
}
