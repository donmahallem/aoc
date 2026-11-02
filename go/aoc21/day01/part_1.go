package day01

import (
	"bufio"
	"io"
	"strconv"
)

func Part1(in io.Reader) (int, error) {
	s := bufio.NewScanner(in)
	increased := 0
	summe := 0
	previous := -1
	for s.Scan() {
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			return 0, err
		}
		if previous >= 0 && val > previous {
			increased++
		}
		previous = val
		summe += val
	}
	if s.Err() != nil {
		return 0, s.Err()
	}
	return increased, nil
}
