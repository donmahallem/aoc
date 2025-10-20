package day12

import (
	"io"
)

func Part2(in io.Reader) int {
	lines := ParseInput(in, 5)
	var total int = 0
	for _, line := range lines {
		total += SolveLine(line)
	}
	return total
}
