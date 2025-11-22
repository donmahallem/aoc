package day24

import (
	"io"
)

func Part1(in io.Reader) int {
	inp := parseInput[float64](in)
	return findCollisions(inp, 200000000000000, 400000000000000)
}
