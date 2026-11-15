package day01

import "io"

func Part1(in io.Reader) (int, error) {
	nums, err := parse(in)
	if err != nil {
		return 0, err
	}
	maxValue := 0
	for _, n := range nums {
		if n > maxValue {
			maxValue = n
		}
	}
	return maxValue, nil
}
