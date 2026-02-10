package day10

import (
	"io"
)

func loopLengthFrom(start *node) int {
	if start == nil {
		return 0
	}
	// Pick one of the two neighbors of S
	var cur *node
	for _, nb := range []*node{start.Up, start.Right, start.Down, start.Left} {
		if nb != nil {
			cur = nb
			break
		}
	}
	if cur == nil {
		return 0
	}
	prev := start
	length := 1
	for cur != start {
		var next *node
		for _, nb := range []*node{cur.Up, cur.Right, cur.Down, cur.Left} {
			if nb != nil && nb != prev {
				next = nb
				break
			}
		}
		if next == nil {
			return 0 // not a loop
		}
		prev, cur = cur, next
		length++
	}
	return length
}

func Part1(in io.Reader) (int, error) {
	start, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	loopSize := loopLengthFrom(start)
	return loopSize / 2, nil
}
