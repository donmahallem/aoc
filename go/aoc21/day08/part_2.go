package day08

import (
	"io"
)

func Part2(in io.Reader) (int, error) {
	counts, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	summe := 0
	for lineIdx := range counts {
		decoded := decodeMapping(counts[lineIdx].mapping)
		current := 0
		for _, out := range counts[lineIdx].output {
			current = current*10 + decoded[out]
		}
		summe += current
	}
	return summe, nil
}
