package day05

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func FixLine(facts *map[int][]int, line *[]int) (int, bool) {
	slices.SortFunc(*line, func(i, j int) int {
		if slices.Contains((*facts)[j], i) {
			return 1
		}
		return -1
	})
	if val, ok := ValidateLine(facts, line); ok {
		return val, true
	}
	return -1, false
}
func Part2(in *os.File) {
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
			if _, ok := ValidateLine(&m, &parsedLine); ok {
				continue //counter += midValue
			} else if midValue, ok := FixLine(&m, &parsedLine); ok {
				counter += midValue
			}
		}
	}
	fmt.Printf("%d\n", counter)
}
