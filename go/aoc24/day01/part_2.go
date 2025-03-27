package day01

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func count[T int](slice []T, val T) int {
	count := 0
	for _, s := range slice {
		if s == val {
			count++
		}
	}
	return count
}

func Part2(in io.Reader) {
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
	for i := 0; i < len(left); i++ {
		summe += left[i] * count(right, left[i])
	}
	fmt.Printf("Result: %d\n", summe)
}
