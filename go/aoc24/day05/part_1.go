package day05

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
	"github.com/donmahallem/aoc/go/aoc_utils/bytes"
)

type facts map[int64]struct{}

func encodePair(a, b int64) int64 {
	return (a << 32) | b
}

// validateLine ensures no ordering rules are violated and returns the median entry.
func validateLine(f facts, line []int64) (int64, bool) {
	n := len(line)
	if n == 0 {
		return 0, false
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if _, violates := f[encodePair(line[j], line[i])]; violates {
				return 0, false
			}
		}
	}
	return line[n/2], true
}

func Part1(in io.Reader) (int64, error) {
	s := bufio.NewScanner(in)
	baseData := true
	counter := int64(0)
	var numberA, numberB int64
	encodedRules := make(facts)
	numbers := make([]int64, 0, 16)

	for s.Scan() {
		lineData := s.Bytes()
		if len(lineData) == 0 {
			baseData = false
			continue
		}
		if baseData {
			numberA, numberB = 0, 0
			currentNumber := &numberA
			for _, c := range lineData {
				if bytes.ByteIsNumber(c) {
					*currentNumber = (*currentNumber)*10 + int64(c-'0')
				} else if c == '|' {
					currentNumber = &numberB
				} else {
					return 0, aoc_utils.NewUnexpectedInputError(c)
				}
			}
			encodedRules[encodePair(numberA, numberB)] = struct{}{}
			continue
		}

		bytes.ParseIntSequence(lineData, ',', &numbers)
		if len(numbers) == 0 {
			continue
		}
		if midValue, ok := validateLine(encodedRules, numbers); ok {
			counter += midValue
		}
	}
	return counter, nil
}
