package day01

import (
	"bufio"
	"io"
	"slices"
)

var da = [9][]byte{[]byte("one"),
	[]byte("two"),
	[]byte("three"),
	[]byte("four"),
	[]byte("five"),
	[]byte("six"),
	[]byte("seven"),
	[]byte("eight"),
	[]byte("nine")}

func ParseLinePart2(line []byte) int {
	var curVal, lastVal uint8 = 20, 20
	for i := range line {
		if (line)[i] >= '0' && (line)[i] <= '9' {
			if curVal >= 10 {
				curVal = (line)[i] - '0'
				lastVal = (line)[i] - '0'
			} else {
				lastVal = (line)[i] - '0'
			}
		} else {
			for j := range da {
				exp := line[i : i+len(da[j])]
				if slices.Equal(exp, da[j]) {
					if curVal >= 10 {
						curVal = uint8(j + 1)
						lastVal = uint8(j + 1)
					} else {
						lastVal = uint8(j + 1)
					}
				}
			}
		}
	}
	return int(curVal*10 + lastVal)
}
func Part2(in io.Reader) int {
	s := bufio.NewScanner(in)
	summe := 0
	for s.Scan() {
		summe += ParseLinePart2(s.Bytes())
	}
	return summe
}
