package day09

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type InputRow []int
type Input []InputRow

func ParseInput(in io.Reader) Input {
	s := bufio.NewScanner(in)
	result := make(Input, 0)
	for s.Scan() {
		parts := strings.Fields(s.Text())
		nums := make([]int, len(parts))
		for idx, item := range parts {
			nums[idx], _ = strconv.Atoi(item)
		}
		result = append(result, nums)
	}
	return result
}

func PredictRight(row InputRow) int {
	for endIdx := len(row); endIdx > 0; endIdx-- {
		allZero := true
		for idx := range endIdx - 1 {
			row[idx] = row[idx+1] - row[idx]
			if allZero && row[idx] != 0 {
				allZero = false
			}
		}
		if allZero {
			cumSum := 0
			for upIdx := endIdx - 1; upIdx < len(row); upIdx++ {
				cumSum += row[upIdx]
			}
			return cumSum
		}
	}
	return 0
}

func Part1(in io.Reader) int {
	rows := ParseInput(in)
	cumSum := 0
	for _, row := range rows {
		cumSum += PredictRight(row)
	}
	return cumSum
}
