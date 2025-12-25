package day23

import (
	"io"
)

func Part2(in io.Reader) (int, error) {
	inp, w, h, err := parseInput(in, false)
	if err != nil {
		return 0, err
	}
	g, startIdx, endIdx := createGraph(inp, w, h)
	return dfsIterative(g, startIdx, endIdx), nil
}
