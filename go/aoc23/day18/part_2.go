package day18

import (
	"io"
)

func Part2(in io.Reader) int64 {
	start := ParseInput(in, false)
	return circumcise(start)
}
