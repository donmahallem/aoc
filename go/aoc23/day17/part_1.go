package day17

import (
	"io"
)

func Part1(r io.Reader) (uint32, error) {
	f, err := parseInput(r)
	if err != nil {
		return 0, err
	}
	return findShortestPath(f, 0, 3), nil
}
