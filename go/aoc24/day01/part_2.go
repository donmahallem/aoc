package day01

import (
	"bufio"
	"io"
	"strconv"
	"strings"
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

func Part2(in io.Reader) int {
	s := bufio.NewScanner(in)
	left := make([]int, 0)
	right := make([]int, 0)
	for s.Scan() {
		var line = strings.Split(s.Text(), "   ")
		var int_left, _ = strconv.Atoi(line[0])
		var int_right, _ = strconv.Atoi(line[1])
		left = append(left, int_left)
		right = append(right, int_right)
	}

	var summe int = 0
	for i := range len(left) {
		summe += left[i] * Count(right, left[i])
	}
	return summe
}
