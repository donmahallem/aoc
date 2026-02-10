package day21

import (
	"io"
)

func Part2(in io.Reader) (uint, error) {
	return CalculateMoves(in, 26), nil
}

//306335137543664
//49502460620748
