package day17

import "io"

func Part2(in io.Reader) uint32 {
	start := ParseInput(in)
	return findShortestPath(start, 4, 10)
}
