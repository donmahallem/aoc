package day23

import (
	"io"
)

func Part1(in io.Reader) int {
	inp, w, h := parseInput(in, true)
	g, startIdx, endIdx := createGraph(inp, w, h)
	return dfsIterative(g, startIdx, endIdx)
}
