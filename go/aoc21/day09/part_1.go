package day09

import (
	"io"
)

func Part1(in io.Reader) (int, error) {
	data, width, height, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	count := 0
	for y := range height {
		for x := range width {
			idx := y*width + x
			value := data[idx]
			if x > 0 && data[idx-1] <= value {
				continue
			}
			if x < width-1 && data[idx+1] <= value {
				continue
			}
			if y > 0 && data[idx-width] <= value {
				continue
			}
			if y < height-1 && data[idx+width] <= value {
				continue
			}
			count += int(value-'0') + 1
		}
	}
	return count, nil
}
