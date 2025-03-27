package day11

import (
	"io"
)

func Part2(in io.Reader) int {
	data, _ := ParseLine(in)

	return SplitStones(data, 75)
}
