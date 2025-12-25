package day13

import (
	"io"
)

func Part2(in io.Reader) (int, error) {
	data, err := LoadFile(in)
	if err != nil {
		return 0, err
	}
	totalSum := 0
	for _, inp := range data {
		inp.target.Y += 10000000000000
		inp.target.X += 10000000000000
		a, b, ok := Calculate(&inp)
		if !ok {
			continue
		}
		totalSum += a*3 + b
	}
	return totalSum, nil
}
