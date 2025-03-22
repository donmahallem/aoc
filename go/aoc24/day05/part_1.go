package day05

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ValidateLine(facts *map[int][]int, line *[]int) (int, bool) {
	for i := 1; i < len(*line); i++ {
		if !slices.Contains((*facts)[(*line)[i-1]], (*line)[i]) {
			return -1, false
		}
	}
	return (*line)[len(*line)/2], true
}

func ParseLine(line string) ([]int, error) {
	data := strings.Split(line, ",")
	parsedData := make([]int, len(data))
	for idx, dataStr := range data {
		currentNum, err := strconv.Atoi(dataStr)
		if err != nil {
			return nil, err
		}
		parsedData[idx] = currentNum
	}
	return parsedData, nil
}

func Part1(in *os.File) {
	s := bufio.NewScanner(in)
	m := make(map[int][]int)
	baseData := true
	counter := 0
	for s.Scan() {
		lineData := s.Text()
		if len(lineData) == 0 {
			baseData = false
		} else if baseData {
			data := strings.Split(lineData, "|")
			num_a, _ := strconv.Atoi(data[0])
			num_b, _ := strconv.Atoi(data[1])
			if _, ok := m[num_a]; ok {
				m[num_a] = append(m[num_a], num_b)
			} else {
				m[num_a] = []int{num_b}
			}
		} else {
			parsedLine, _ := ParseLine(lineData)
			if midValue, ok := ValidateLine(&m, &parsedLine); ok {
				counter += midValue
			}
		}
	}
	fmt.Printf("%d\n", counter)
}
