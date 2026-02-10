package day14

import (
	"io"
)

func Part1(in io.Reader) (uint, error) {
	start, err := parseInputPart1(in)
	if err != nil {
		return 0, err
	}
	accum := uint(0)
	for idx, val := range start {
		accum += val * uint(len(start)-idx)
	}
	return accum, nil
}
