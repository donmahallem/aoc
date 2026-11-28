package day09

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func parseInput(inputData io.Reader) (data []byte, width int, height int, err error) {
	scanner := bufio.NewScanner(inputData)
	expectedLineLength := -1
	rows := 0

	for scanner.Scan() {
		line := scanner.Bytes()
		if expectedLineLength == -1 {
			expectedLineLength = len(line)
			data = make([]byte, expectedLineLength*expectedLineLength) // assuming square grid for simplicity
		} else if len(line) != expectedLineLength {
			return nil, 0, 0, aoc_utils.NewParseErrorf("invalid line length: %s", line)
		}
		copy(data[rows*expectedLineLength:], line)
		rows++
	}
	if err := scanner.Err(); err != nil {
		return nil, 0, 0, err
	}
	return data, expectedLineLength, rows, nil
}
