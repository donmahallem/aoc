package day01

import (
	"bufio"
	"io"
)

func ParseLine(line []byte) int {
	var curVal, lastVal uint8 = 20, 20
	for i := range line {
		if (line)[i] >= '0' && (line)[i] <= '9' {
			if curVal >= 10 {
				curVal = (line)[i] - '0'
				lastVal = (line)[i] - '0'
			} else {
				lastVal = (line)[i] - '0'
			}
		}
	}
	return int(curVal*10 + lastVal)
}
func Part1(in io.Reader) int {
	s := bufio.NewScanner(in)
	summe := 0
	for s.Scan() {
		summe += ParseLine(s.Bytes())
	}
	return summe
}
