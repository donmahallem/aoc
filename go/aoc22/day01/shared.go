package day01

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func parse(in io.Reader) ([]int, error) {
	parsedNumbers := make([]int, 0, 256)
	s := bufio.NewScanner(in)
	currentSum := 0
	for s.Scan() {
		line := s.Bytes()
		if len(line) == 0 {
			parsedNumbers = append(parsedNumbers, currentSum)
			currentSum = 0
			continue
		}
		num := 0
		for _, b := range line {
			if b < '0' || b > '9' {
				return nil, aoc_utils.NewUnexpectedInputError(b)
			}
			num = num*10 + int(b-'0')
		}
		currentSum += num
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	if currentSum > 0 {
		parsedNumbers = append(parsedNumbers, currentSum)
	}
	return parsedNumbers, nil
}
