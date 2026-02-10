package day11

import (
	"io"
)

func Part2(in io.Reader) (int, error) {
	start, err := parseInput(in, 999999)
	if err != nil {
		return 0, err
	}
	return combinedManhattenDistances(start), nil
}
