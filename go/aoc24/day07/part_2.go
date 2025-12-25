package day07

import (
	"bufio"
	"io"
	"math"

	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

func numDigits(val int) int {
	return int(math.Log10(float64(val))) + 1
}

func opConcat(a, b int) int {
	offset := numDigits(b)
	return (a * int_util.IntPow(10, offset)) + b
}

func checkLinePart2(line *parsedLineData) bool {
	numTerms := len(line.TestValues)
	if numTerms == 0 {
		return false
	}
	runnerTarget := int_util.IntPow(3, numTerms-1)
	for i := range runnerTarget {
		testResult := line.TestValues[0]
		operationType := i
		for pos := 1; pos < numTerms; pos++ {
			operation := operationType % 3
			operationType /= 3
			switch operation {
			case 0:
				testResult += line.TestValues[pos]
			case 1:
				testResult *= line.TestValues[pos]
			case 2:
				testResult = opConcat(testResult, line.TestValues[pos])
			}
			if testResult > line.Result {
				break
			}
		}
		if testResult == line.Result {
			return true
		}
	}
	return false
}

func Part2(in io.Reader) (int, error) {
	s := bufio.NewScanner(in)
	validSum := 0
	parsedLine := parsedLineData{}
	for s.Scan() {
		line := s.Bytes()
		parseLine(line, &parsedLine)
		if checkLinePart2(&parsedLine) {
			validSum += parsedLine.Result
		}
	}
	return validSum, nil
}
