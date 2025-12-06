package day06

import (
	"bufio"
	"io"
)

func add(a, b uint64) uint64 {
	return a + b
}

func mul(a, b uint64) uint64 {
	return a * b
}

type operator func(a, b uint64) uint64

func Part2(in io.Reader) uint64 {
	var columns []uint64 = nil
	scanner := bufio.NewScanner(in)
	totalNumber := uint64(0)
	var currentOperator operator // zero value is nil
	for scanner.Scan() {
		line := scanner.Bytes()
		if columns == nil {
			columns = make([]uint64, len(line))
		}
		currentNumber := uint64(0)
		for i := range len(line) {
			b := line[i]
			if b >= '0' && b <= '9' {
				columns[i] = (columns[i] * 10) + uint64(b-'0')
			} else if b == '+' {
				totalNumber += currentNumber
				currentNumber = columns[i]
				currentOperator = add
			} else if b == '*' {
				totalNumber += currentNumber
				currentNumber = columns[i]
				currentOperator = mul
			} else if currentOperator != nil && columns[i] != 0 && b == ' ' {
				currentNumber = currentOperator(currentNumber, columns[i])
			}
		}
		// the operator line may be shorter than the number of columns
		if currentOperator != nil {
			for i := len(line); i < len(columns); i++ {
				currentNumber = currentOperator(currentNumber, columns[i])
			}
		}
		totalNumber += currentNumber
	}
	return totalNumber
}
