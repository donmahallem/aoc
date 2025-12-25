package day23

import (
	"io"
)

func Part1(in io.Reader) (int, error) {
	inp, w, h, err := parseInput(in, true)
	if err != nil {
		return 0, err
	}
	g, startIdx, endIdx := createGraph(inp, w, h)
	return dfsIterative(g, startIdx, endIdx), nil
}
