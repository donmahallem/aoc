package day01

import (
	"io"
	"slices"
)

func Part2(in io.Reader) (int, error) {
	nums, err := parse(in)
	if err != nil {
		return 0, err
	}

	slices.SortFunc(nums, func(a, b int) int {
		return b - a
	})
	sum := 0
	for i := 0; i < 3 && i < len(nums); i++ {
		sum += nums[i]
	}
	return sum, nil
}
