package day02

import "io"

func Part2(in io.Reader) int {
	invalidSum := 0

	for val := range parseInputGen(in) {

		invalids := findRepeatedBlocks(val)
		for num, _ := range invalids {
			invalidSum += int(num)
		}
	}
	return invalidSum
}
