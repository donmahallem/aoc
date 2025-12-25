package day15

import (
	"bufio"
	"io"
)

func Part1(in io.Reader) (uint, error) {
	scanner := bufio.NewScanner(in)
	runningTotal := uint(0)
	currentTotal := uint(0)
	for scanner.Scan() {
		line := scanner.Bytes()
		for c := range line {
			switch line[c] {
			case ',':
				runningTotal += currentTotal
				currentTotal = 0
			default:
				currentTotal = (currentTotal + uint(line[c])) * 17 % 256
			}
		}
	}
	if currentTotal > 0 {

		runningTotal += currentTotal
	}
	return runningTotal, nil
}
