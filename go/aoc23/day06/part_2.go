package day06

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/aoc_utils/bytes"
)

func parseInputPart2(in io.Reader) Race {
	var time int
	s := bufio.NewScanner(in)
	firstNumber := true
	currentNumber := 0
	for s.Scan() {
		b := s.Bytes()
		for _, c := range b {
			if parsedInt, ok := bytes.ParseIntFromByte[int](c); ok {
				currentNumber *= 10
				currentNumber += parsedInt
			}
		}
		if firstNumber {
			time = currentNumber
			currentNumber = 0
			firstNumber = false
		} else {
			break
		}
	}
	return Race{Time: time, Distance: currentNumber}
}

func Part2(in io.Reader) int {
	races := parseInputPart2(in)
	return CountOptions(races)
}
