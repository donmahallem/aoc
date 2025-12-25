package day17

import (
	"io"
)

func Part2(r io.Reader) (uint32, error) {
	f, err := parseInput(r)
	if err != nil {
		return 0, err
	}
	return findShortestPath(f, 4, 10), nil
}
