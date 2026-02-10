package day18

import (
	"io"
)

func Part2(r io.Reader) (int64, error) {
	ins, err := parseInput(r, false)
	if err != nil {
		return 0, err
	}
	return calculateArea(ins), nil
}
