package day07

import (
	"bufio"
	"io"
	"math"

	"github.com/donmahallem/aoc/aoc_utils/math/pow"
)

func NumDigits(val int) int {
	return int(math.Log10(float64(val))) + 1
}

func OpConcat(a, b int) int {
	offset := NumDigits(b)
	return (a * pow.IntPow(10, offset)) + b
}

func CheckLinePart2(line *ParsedLineData) bool {
	numTerms := len(line.TestValues)
	runnerTarget := pow.IntPow(3, numTerms-1)
	for i := range runnerTarget {
		testResult := line.TestValues[0]
		for pos := 1; pos < numTerms; pos++ {
			switch (i / pow.IntPow(3, pos-1)) % 3 {
			case 0:
				testResult += line.TestValues[pos]
			case 1:
				testResult *= line.TestValues[pos]
			case 2:
				testResult = OpConcat(testResult, line.TestValues[pos])
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

func Part2(in io.Reader) int {
	s := bufio.NewScanner(in)
	validSum := 0
	parsedLine := ParsedLineData{}
	for s.Scan() {
		line := s.Bytes()
		ParseLine(line, &parsedLine)
		if CheckLinePart2(&parsedLine) {
			validSum += parsedLine.Result
		}
	}
	return validSum
}
