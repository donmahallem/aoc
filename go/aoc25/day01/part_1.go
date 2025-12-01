package day01

import (
	"io"
)

func Part1(in io.Reader) int {
	currentPosition := 50
	zeros := 0
	for d := range parseInputGen(in) {
		currentPosition = (currentPosition + d) % 100
		if currentPosition == 0 {
			zeros++
		}
	}
	return zeros
}
