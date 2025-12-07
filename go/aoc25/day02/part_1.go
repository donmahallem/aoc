package day02

import (
	"io"
)

func Part1(in io.Reader) int {
	invalidSum := 0

	for val := range parseInputGen(in) {

		invalids := findRepeatedBlocks(val)
		for num, k := range invalids {
			if k == 2 {
				invalidSum += int(num)
			}
		}
	}
	return invalidSum
}
