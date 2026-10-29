package day01

import (
	"bufio"
	"io"
	"strconv"
)

func Part2(in io.Reader) (int, error) {
	s := bufio.NewScanner(in)
	var window [4]int
	count := 0
	i := 0

	for s.Scan() {
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			return 0, err
		}

		window[i%4] = val

		// Accumulated enough values to compare
		if i >= 3 {
			prevSumIndex := (i + 1) % 4
			currSumIndex := i % 4

			// We compare (A+B+C)>(B+C+D) <=> A > D
			if window[currSumIndex] > window[prevSumIndex] {
				count++
			}
		}
		i++
	}

	if s.Err() != nil {
		return 0, s.Err()
	}
	return count, nil
}
