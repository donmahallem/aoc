package day07

import (
	"io"
	"slices"
)

func Part1(in io.Reader) (int, error) {
	counts, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	slices.Sort(counts)
	median := counts[len(counts)/2]
	summe := 0
	for _, v := range counts {
		diff := v - median
		if diff < 0 {
			diff = -diff
		}
		summe += diff
	}
	return summe, nil
}
