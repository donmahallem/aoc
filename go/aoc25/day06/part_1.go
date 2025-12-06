package day06

import (
	"bufio"
	"io"
)

type pairs struct {
	mul uint64
	add uint64
}

func Part1(in io.Reader) uint64 {
	var columns []pairs = make([]pairs, 0, 32)
	scanner := bufio.NewScanner(in)
	totalNumber := uint64(0)
	firstRow := true
	for scanner.Scan() {
		line := scanner.Bytes()
		currentNumber := uint64(0)
		currentNumberIdx := 0
		for charIdx := range len(line) {
			b := line[charIdx]
			if currentNumber > 0 && b == ' ' {
				if firstRow {
					columns = append(columns, pairs{mul: currentNumber, add: currentNumber})
				} else {
					columns[currentNumberIdx].mul *= currentNumber
					columns[currentNumberIdx].add += currentNumber
				}
				currentNumber = 0
				currentNumberIdx++
			} else if b >= '0' && b <= '9' {
				currentNumber = currentNumber*10 + uint64(b-'0')
			} else if b == '+' {
				totalNumber += columns[currentNumberIdx].add
				currentNumberIdx++
			} else if b == '*' {
				totalNumber += columns[currentNumberIdx].mul
				currentNumberIdx++
			}
		}
		if currentNumber > 0 {
			if firstRow {
				columns = append(columns, pairs{mul: currentNumber, add: currentNumber})
			} else {
				columns[currentNumberIdx].mul *= currentNumber
				columns[currentNumberIdx].add += currentNumber
			}
		}
		firstRow = false
	}

	return totalNumber
}
