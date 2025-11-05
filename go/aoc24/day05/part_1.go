package day05

import (
	"bufio"
	"fmt"
	"io"

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

// parseLine tokenises a comma-separated list of page numbers into dst.
func parseLine(line []byte, dst []int64) []int64 {
	dst = dst[:0]
	var value int64
	hasDigit := false
	for _, b := range line {
		if bytes.ByteIsNumber(b) {
			value = value*10 + int64(b-'0')
			hasDigit = true
			continue
		}
		if hasDigit {
			dst = append(dst, value)
			value = 0
			hasDigit = false
		}
	}
	if hasDigit {
		dst = append(dst, value)
	}
	return dst
}

func Part1(in io.Reader) int64 {
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
					panic(fmt.Sprintf("Unexpected character %c in base data", c))
				}
			}
			encodedRules[encodePair(numberA, numberB)] = struct{}{}
			continue
		}

		numbers = parseLine(lineData, numbers)
		if len(numbers) == 0 {
			continue
		}
		if midValue, ok := validateLine(encodedRules, numbers); ok {
			counter += midValue
		}
	}
	return counter
}
