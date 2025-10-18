package day11

import (
	"io"
)

func Part2(in io.Reader) int {
	start := ParseInput(in, 999999)
	return combinedManhattenDistances(start)
}
