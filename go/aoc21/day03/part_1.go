package day03

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func Part1(in io.Reader) (int, error) {
	s := bufio.NewScanner(in)
	if !s.Scan() {
		if s.Err() != nil {
			return 0, s.Err()
		}
		return 0, aoc_utils.NewParseErrorf("empty input")
	}
	firstLine := s.Bytes()
	expectedLineLength := len(firstLine)
	lines := 1
	bitBuffer := make([]int, expectedLineLength)

	accumulateBits := func(line []byte) error {
		for i, b := range line {
			if b == '1' {
				bitBuffer[i]++
			} else if b != '0' {
				return aoc_utils.NewParseErrorf("invalid character in input: %c", b)
			}
		}
		return nil
	}
	if err := accumulateBits(firstLine); err != nil {
		return 0, err
	}
	for s.Scan() {
		line := s.Bytes()
		if len(line) != expectedLineLength {
			return 0, aoc_utils.NewParseErrorf("unexpected line length: got %d, expected %d", len(line), expectedLineLength)
		}
		if err := accumulateBits(line); err != nil {
			return 0, err
		}
		lines++
	}

	if s.Err() != nil {
		return 0, s.Err()
	}
	gammaRate, epsilonRate := 0, 0
	for i, b := range bitBuffer {
		if b*2 >= lines {
			gammaRate |= 1 << (expectedLineLength - 1 - i)
		} else {
			epsilonRate |= 1 << (expectedLineLength - 1 - i)
		}
	}
	return gammaRate * epsilonRate, nil
}
