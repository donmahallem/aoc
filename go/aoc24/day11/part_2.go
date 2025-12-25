package day11

import (
	"io"
)

func Part2(in io.Reader) (int, error) {
	data, err := parseLine(in)
	if err != nil {
		return 0, err
	}
	return splitStones(data, 75), nil
}
