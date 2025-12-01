package day01

import (
	"io"
)

func mod100(n int) int {
	n %= 100
	if n < 0 {
		n += 100
	}
	return n
}

func Part2(in io.Reader) int {
	data := parseInput(in)
	currentPosition := 50
	zeros := 0
	for _, d := range data {
		start := currentPosition
		if d.distance == 0 {
			if start == 0 {
				zeros++
			}
			// position unchanged
			continue
		}

		zeroHits := 0
		if d.left {
			zeroHits = start % 100
		} else {
			zeroHits = (100 - start) % 100
		}
		if zeroHits == 0 {
			zeroHits = 100
		}
		if d.distance >= zeroHits {
			zeros += 1 + (d.distance-zeroHits)/100
		}

		if d.left {
			currentPosition = mod100(start - d.distance)
		} else {
			currentPosition = mod100(start + d.distance)
		}
	}
	return zeros
}
