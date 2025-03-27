package day07

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"sync"
)

func CheckLinePart1(result *int, terms *[]int) bool {
	numTerms := len(*terms)
	runnerTarget := (1 << (numTerms - 1))
	for i := range runnerTarget {
		testResult := (*terms)[0]
		for pos := 1; pos < numTerms; pos++ {
			if (1<<(pos-1))&i > 0 {
				testResult += (*terms)[pos]
			} else {
				testResult *= (*terms)[pos]
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

func parseLine(lineData *string) (*int, *[]int) {
	parts := strings.Split(*lineData, ": ")
	expectedSum, _ := strconv.Atoi(parts[0])
	items := strings.Split(parts[1], " ")
	terms := make([]int, len(items))
	for idx, item := range items {
		terms[idx], _ = strconv.Atoi(item)
	}
	return &expectedSum, &terms
}

func Part1(in io.Reader) int {
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
			if CheckLinePart1(expectedSum, terms) {
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
