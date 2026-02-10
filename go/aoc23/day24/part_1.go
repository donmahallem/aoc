package day24

import (
	"io"
)

func Part1(in io.Reader) (int, error) {
	inp, err := parseInput[float64](in)
	if err != nil {
		return 0, err
	}
	return findCollisions(inp, 200000000000000, 400000000000000), nil
}
