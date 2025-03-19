package day05

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ValidateLine(facts *map[int][]int, line string) (int, bool) {
	data := strings.Split(line, ",")
	lastNum, _ := strconv.Atoi(data[0])
	midIdx := len(data) / 2
	midValue := -1
	for i := 1; i < len(data); i++ {
		currentNum, _ := strconv.Atoi(data[i])
		if !slices.Contains((*facts)[lastNum], currentNum) {
			return -1, false
		}
		if midIdx == i {
			midValue = currentNum
		}
		lastNum = currentNum
	}
	return midValue, true
}

func Part1() {
	s := bufio.NewScanner(os.Stdin)
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
			if midValue, ok := ValidateLine(&m, lineData); ok {
				counter += midValue
			}
		}
	}
	fmt.Printf("%d\n", counter)
}
