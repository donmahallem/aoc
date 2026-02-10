package day18

import (
	"io"
)

func Part1(r io.Reader) (int64, error) {
	ins, err := parseInput(r, true)
	if err != nil {
		return 0, err
	}
	return calculateArea(ins), nil
}
