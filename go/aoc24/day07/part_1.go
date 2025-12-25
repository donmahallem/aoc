package day07

import (
	"bufio"
	"io"
)

func checkLinePart1(parsed *parsedLineData) bool {
	numTerms := len(parsed.TestValues)
	if numTerms == 0 {
		return false
	}
	// limit large inputs from fuzzing to avoid combinatorial blowup
	if numTerms > 20 {
		return false
	}
	runnerTarget := 1 << (numTerms - 1)
	for i := 0; i < runnerTarget; i++ {
		testResult := parsed.TestValues[0]
		for pos := 1; pos < numTerms; pos++ {
			if (1<<(pos-1))&i > 0 {
				testResult += parsed.TestValues[pos]
			} else {
				testResult *= parsed.TestValues[pos]
			}
			if testResult > parsed.Result {
				break
			}
		}
		if testResult == parsed.Result {
			return true
		}
	}
	return false
}

func Part1(in io.Reader) (int, error) {
	s := bufio.NewScanner(in)
	validSum := 0
	parsedLine := parsedLineData{}
	for s.Scan() {
		line := s.Bytes()
		parseLine(line, &parsedLine)
		if checkLinePart1(&parsedLine) {
			validSum += parsedLine.Result
		}
	}
	return validSum, nil
}
