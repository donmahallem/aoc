package day21

import (
	"io"
)

func Part2(in io.Reader) int {
	inp := ParseInput(in)
	return CountVisited(&inp, 26501365)
}
