package day01

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func Count[T int](slice []T, val T) int {
	count := 0
	for _, s := range slice {
		if s == val {
			count++
		}
	}
	return count
}

func Part2(in io.Reader) (any, error) {
	s := bufio.NewScanner(in)
	left := make([]int, 0)
	right := make([]int, 0)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		if len(fields) < 2 {
			return nil, aoc_utils.NewUnexpectedInputError(0)
		}
		int_left, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, aoc_utils.NewParseError("invalid left value", err)
		}
		int_right, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, aoc_utils.NewParseError("invalid right value", err)
		}
		left = append(left, int_left)
		right = append(right, int_right)
	}

	var summe int = 0
	for i := range left {
		summe += left[i] * Count(right, left[i])
	}
	return summe, nil
}
