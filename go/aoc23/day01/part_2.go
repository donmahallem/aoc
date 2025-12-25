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

func parseLinePart2(line []byte) (int, error) {
	var curVal, lastVal int = 20, 20
	for i := range line {
		if (line)[i] >= '0' && (line)[i] <= '9' {
			if curVal >= 10 {
				curVal = int((line)[i] - '0')
				lastVal = int((line)[i] - '0')
			} else {
				lastVal = int((line)[i] - '0')
			}
		} else {
			for searchSequenceIdx := range da {
				// ensure we don't slice past the end of the line
				if i+len(da[searchSequenceIdx]) <= len(line) {
					exp := line[i : i+len(da[searchSequenceIdx])]
					if slices.Equal(exp, da[searchSequenceIdx]) {
						if curVal >= 10 {
							curVal = searchSequenceIdx + 1
							lastVal = searchSequenceIdx + 1
						} else {
							lastVal = searchSequenceIdx + 1
						}
					}
				}
			}
		}
	}
	return curVal*10 + lastVal, nil
}
func Part2(in io.Reader) (int, error) {
	s := bufio.NewScanner(in)
	summe := 0
	for s.Scan() {
		val, err := parseLinePart2(s.Bytes())
		if err != nil {
			return 0, err
		}
		summe += val
	}
	return summe, nil
}
