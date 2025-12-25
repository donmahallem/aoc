package day12

import (
	"io"
)

func Part2(in io.Reader) (int, error) {
	lines, err := parseInput(in, 5)
	if err != nil {
		return 0, err
	}
	var total int = 0
	for _, line := range lines {
		stepCount, err := solveLine(*line)
		if err != nil {
			return 0, err
		}
		total += stepCount
	}
	return total, nil
}
