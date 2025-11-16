package day23

import (
	"io"
)

func Part2(in io.Reader) int {
	inp, w, h := parseInput(in, false)
	g, startIdx, endIdx := createGraph(inp, w, h)
	return dfsIterative(g, startIdx, endIdx)
}
