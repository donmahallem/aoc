package day07

import (
	"bufio"
	"io"
	"math"
	"sync"

	"github.com/donmahallem/aoc/aoc_utils"
)

func NumDigits(val int) int {
	return int(math.Log10(float64(val))) + 1
}

func OpConcat(a, b int) int {
	offset := NumDigits(b)
	return (a * aoc_utils.IntPow(10, offset)) + b
}

func CheckLinePart2(result *int, terms *[]int) bool {
	numTerms := len(*terms)
	runnerTarget := aoc_utils.IntPow(3, numTerms-1)
	for i := range runnerTarget {
		testResult := (*terms)[0]
		for pos := 1; pos < numTerms; pos++ {
			switch (i / aoc_utils.IntPow(3, pos-1)) % 3 {
			case 0:
				testResult += (*terms)[pos]
				break
			case 1:
				testResult *= (*terms)[pos]
				break
			case 2:
				testResult = OpConcat(testResult, (*terms)[pos])
				break
			}
			if testResult > *result {
				break
			}
		}
		if testResult == *result {
			return true
		}
	}
	return false
}

func Part2(in io.Reader) int {
	s := bufio.NewScanner(in)
	validSum := 0
	resultChannel := make(chan int, 50000)
	var wg sync.WaitGroup
	for s.Scan() {
		line := s.Text()
		wg.Add(1)
		go func(lineData *string) {
			defer wg.Done()
			expectedSum, terms := parseLine(lineData)
			if CheckLinePart2(expectedSum, terms) {
				resultChannel <- *expectedSum
			}
		}(&line)
	}
	wg.Wait()
	close(resultChannel)
	for item := range resultChannel {
		validSum += item
	}
	return validSum
}
