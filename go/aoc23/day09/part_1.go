package day09

import (
	"io"
)

func predictRight(row inputRow) int {
	for endIdx := len(row); endIdx > 0; endIdx-- {
		allZero := true
		for idx := range endIdx - 1 {
			row[idx] = row[idx+1] - row[idx]
			if allZero && row[idx] != 0 {
				allZero = false
			}
		}
		if allZero {
			cumSum := 0
			for upIdx := endIdx - 1; upIdx < len(row); upIdx++ {
				cumSum += row[upIdx]
			}
			return cumSum
		}
	}
	return 0
}

func Part1(in io.Reader) (int, error) {
	rows, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	cumSum := 0
	for _, row := range rows {
		cumSum += predictRight(row)
	}
	return cumSum, nil
}
