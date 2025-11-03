package day01

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils/bytes"
)

func ParseLine(line []byte) int {
	var curVal, lastVal uint8 = 20, 20
	for i := range line {
		if val, ok := bytes.ParseIntFromByte[uint8](line[i]); ok {
			if curVal >= 10 {
				curVal = val
				lastVal = val
			} else {
				lastVal = val
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
