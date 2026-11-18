package day03

import (
	"bufio"
	"io"
)

const sequenceLength = 12

func Part2(in io.Reader) (int, error) {
	accumulator := 0
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		row := scanner.Bytes()

		rowLen := len(row)
		if rowLen <= sequenceLength {
			continue
		}
		start := 0
		currentNumber := 0

		for remaining := sequenceLength; remaining > 0; remaining-- {
			window := row[start : rowLen-remaining+1]
			maxDigit := byte(0)
			maxDigitIndex := 0
			for i, v := range window {
				if v > maxDigit {
					maxDigit = v
					maxDigitIndex = i
				}
			}
			start += maxDigitIndex + 1
			currentNumber = currentNumber*10 + int(maxDigit-baseOffset)
		}
		accumulator += currentNumber
	}
	if scanner.Err() != nil {
		return 0, scanner.Err()
	}
	return accumulator, nil
}
