package day11

import (
	"io"
)

func Part1(in io.Reader) (uint64, error) {
	g, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	counter, err := newPathCounter(g)
	if err != nil {
		return 0, err
	}

	youKey := encodeKey('y', 'o', 'u')
	outKey := encodeKey('o', 'u', 't')
	return counter.count(youKey, outKey), nil
}
