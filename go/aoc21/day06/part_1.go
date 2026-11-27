package day06

import (
	"io"
)

func Part1(in io.Reader) (int, error) {
	counts, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	maxDepth := 80
	totalDevisions := caclcuateDivisions(counts, maxDepth)

	return totalDevisions, nil
}
