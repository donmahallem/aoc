package day21

import (
	"io"
)

func Part2(in io.Reader) uint {
	return CalculateMoves(in, 25)
}
