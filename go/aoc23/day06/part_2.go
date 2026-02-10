package day06

import (
	"io"
)

func Part2(in io.Reader) (int, error) {
	races := parseInputPart2(in)
	return countOptions(races), nil
}
