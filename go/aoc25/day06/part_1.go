package day06

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type pairs struct {
	mul uint64
	add uint64
}

func Part1(in io.Reader) (uint64, error) {
	var columns []pairs = make([]pairs, 0, 32)
	scanner := bufio.NewScanner(in)
	totalNumber := uint64(0)
	firstRow := true
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		numberParsed := false
		currentNumber := uint64(0)
		currentNumberIdx := 0
		for charIdx := range len(line) {
			b := line[charIdx]
			if numberParsed && b == ' ' {
				if firstRow {
					columns = append(columns, pairs{mul: currentNumber, add: currentNumber})
				} else if len(columns) <= currentNumberIdx {
					return 0, aoc_utils.NewParseError("Recieved more numbers than operators", nil)
				} else {
					columns[currentNumberIdx].mul *= currentNumber
					columns[currentNumberIdx].add += currentNumber
				}
				numberParsed = false
				currentNumber = 0
				currentNumberIdx++
			} else if b >= '0' && b <= '9' {
				currentNumber = currentNumber*10 + uint64(b-'0')
				numberParsed = true
			} else if b == '+' {
				if currentNumberIdx >= len(columns) {
					return 0, aoc_utils.NewParseError("Recieved more operators than numbers", nil)
				}
				totalNumber += columns[currentNumberIdx].add
				currentNumberIdx++
			} else if b == '*' {
				if currentNumberIdx >= len(columns) {
					return 0, aoc_utils.NewParseError("Recieved more operators than numbers", nil)
				}
				totalNumber += columns[currentNumberIdx].mul
				currentNumberIdx++
			} else if b != ' ' {
				return 0, aoc_utils.NewUnexpectedInputError(b)
			}
		}
		if numberParsed {
			if firstRow {
				columns = append(columns, pairs{mul: currentNumber, add: currentNumber})
			} else {
				if len(columns) <= currentNumberIdx {
					return 0, aoc_utils.NewParseError("Recieved more numbers than operators", nil)
				}
				columns[currentNumberIdx].mul *= currentNumber
				columns[currentNumberIdx].add += currentNumber
			}
		}
		if len(columns) > 0 {
			firstRow = false
		}
	}

	return totalNumber, nil
}
