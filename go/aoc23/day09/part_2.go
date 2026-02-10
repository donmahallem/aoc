package day09

import (
	"io"
)

func predictLeft(row inputRow) int {
	for startIdx := range len(row) - 1 {
		allZero := true
		for idx := len(row) - 1; idx > startIdx; idx-- {
			row[idx] = row[idx] - row[idx-1]
			if allZero && row[idx] != 0 {
				allZero = false
			}
		}
		if allZero {
			currentValue := 0
			for upIdx := startIdx; upIdx >= 0; upIdx-- {
				currentValue = row[upIdx] - currentValue
			}
			return currentValue
		}
	}
	return 0
}

func Part2(in io.Reader) (int, error) {
	rows, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	cumSum := 0
	for _, row := range rows {
		cumSum += predictLeft(row)
	}
	return cumSum, nil
}
