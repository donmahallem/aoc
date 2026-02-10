package day01

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

func mod100(n int) int {
	n %= 100
	if n < 0 {
		n += 100
	}
	return n
}

func Part2(in io.Reader) (int, error) {
	currentPosition := 50
	zeros := 0
	for d := range parseInputGen(in) {
		start := currentPosition
		if d == 0 {
			if start == 0 {
				zeros++
			}
			continue
		}

		steps := int_util.AbsInt(d)
		left := d < 0

		firstHit := 0
		if left {
			firstHit = start % 100
		} else {
			firstHit = (100 - start) % 100
		}
		if firstHit == 0 {
			firstHit = 100
		}

		if steps >= firstHit {
			zeros += 1 + (steps-firstHit)/100
		}

		if left {
			currentPosition = mod100(start - steps)
		} else {
			currentPosition = mod100(start + steps)
		}
	}
	return zeros, nil
}
