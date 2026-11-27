package day07

import (
	"io"

	"slices"
)

func gauss(n int) int {
	return n * (n + 1) / 2
}

func Part2(in io.Reader) (int, error) {
	counts, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	slices.Sort(counts)
	leftIndex := counts[0]
	rightIndex := counts[len(counts)-1]
	midIndex := (leftIndex + rightIndex) / 2

	cache := make(map[int]int, 64)
	calculateCost := func(pos int) int {
		if cost, ok := cache[pos]; ok {
			return cost
		}
		cost := 0
		for _, v := range counts {
			diff := v - pos
			if diff < 0 {
				diff = -diff
			}
			cost += gauss(diff)
		}
		cache[pos] = cost
		return cost
	}

	for {
		// bisect search
		leftCost := calculateCost(midIndex - 1)
		midCost := calculateCost(midIndex)
		rightCost := calculateCost(midIndex + 1)

		if midCost <= leftCost && midCost <= rightCost {
			return midCost, nil
		}
		if leftCost < midCost {
			rightIndex = midIndex
			midIndex = (leftIndex + rightIndex) / 2
		} else {
			leftIndex = midIndex
			midIndex = (leftIndex + rightIndex) / 2
		}
	}
}
