package day06

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func add(a, b uint64) uint64 {
	return a + b
}

func mul(a, b uint64) uint64 {
	return a * b
}

type operator func(a, b uint64) uint64

func Part2(in io.Reader) (uint64, error) {
	var columns []uint64 = nil
	scanner := bufio.NewScanner(in)
	totalNumber := uint64(0)
	var currentOperator operator // zero value is nil
	expectedLineLength := -1
	for scanner.Scan() {
		line := scanner.Bytes()
		if columns == nil {
			columns = make([]uint64, len(line))
			expectedLineLength = len(line)
		}
		currentNumber := uint64(0)
		if len(line) > expectedLineLength {
			return 0, aoc_utils.NewParseError("inconsistent line lengths in input", nil)
		}
		for i := range len(line) {
			b := line[i]
			if b >= '0' && b <= '9' {
				if len(columns) <= i {
					columns = append(columns, 0)
				}
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
			} else if b != ' ' {
				return 0, aoc_utils.NewUnexpectedInputError(b)
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
	return totalNumber, nil
}
