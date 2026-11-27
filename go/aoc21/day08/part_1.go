package day08

import (
	"io"
)

func Part1(in io.Reader) (int, error) {
	counts, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	count := 0
	for lineIdx := range counts {
		decoded := decodeMapping(counts[lineIdx].mapping)
		for _, out := range counts[lineIdx].output {
			switch decoded[out] {
			case 1, 4, 7, 8:
				count++
			}
		}
	}
	return count, nil
}
