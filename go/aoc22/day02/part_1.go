package day02

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func Part1(in io.Reader) (int, error) {
	parsedNumbers := make([]int, 0, 256)
	s := bufio.NewScanner(in)
	currentSum := 0
	for s.Scan() {
		line := s.Bytes()
		if len(line) == 0 {
			continue
		}
		if len(line) < 3 {
			return 0, aoc_utils.NewParseErrorf("line '%s' is to short", line)
		}
		t1, err := mapToken(line[0])
		if err != nil {
			return 0, err
		}
		t2, err := mapToken(line[2])
		if err != nil {
			return 0, err
		}
		if t2-t1 == 1 || t2-t1 == -2 {
			currentSum += 6
		} else if t1 == t2 {
			currentSum += 3
		}
		currentSum += t2
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	if currentSum > 0 {
		parsedNumbers = append(parsedNumbers, currentSum)
	}
	return currentSum, nil
}
