package day07

import (
	"bufio"
	"io"
)

type ParsedLineData struct {
	Result     int
	TestValues []int
}

func CheckLinePart1(parsed *ParsedLineData) bool {
	numTerms := len(parsed.TestValues)
	runnerTarget := (1 << (numTerms - 1))
	for i := range runnerTarget {
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

func ParseLine(raw []byte, parsed *ParsedLineData) {
	parsed.Result = 0
	parsed.TestValues = parsed.TestValues[:0]
	initialBlock := byte(0)
	currentTestValue := 0
	valueFound := false
	for _, c := range raw {
		if c >= '0' && c <= '9' {
			if initialBlock == 0 {
				parsed.Result = parsed.Result*10 + int(c-'0')
			} else {
				currentTestValue = currentTestValue*10 + int(c-'0')
			}
			valueFound = true
		} else if c == ':' {
			initialBlock = 1
		} else if c == ' ' {
			if initialBlock == 3 {
				parsed.TestValues = append(parsed.TestValues, currentTestValue)
				currentTestValue = 0
				valueFound = false
			} else {
				initialBlock = (initialBlock + 1) | 1
			}
		}
	}
	if valueFound {
		parsed.TestValues = append(parsed.TestValues, currentTestValue)
	}
}

func Part1(in io.Reader) int {
	s := bufio.NewScanner(in)
	validSum := 0
	parsedLine := ParsedLineData{}
	for s.Scan() {
		line := s.Bytes()
		ParseLine(line, &parsedLine)
		if CheckLinePart1(&parsedLine) {
			validSum += parsedLine.Result
		}
	}
	return validSum
}
