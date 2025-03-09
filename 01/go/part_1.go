package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Abs[T int](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	s := bufio.NewScanner(os.Stdin)
	left := make([]int, 0)
	right := make([]int, 0)
	for s.Scan() {
		var line = strings.Split(s.Text(), "   ")
		var int_left, _ = strconv.Atoi(line[0])
		var int_right, _ = strconv.Atoi(line[1])
		left = append(left, int_left)
		right = append(right, int_right)
	}
	slices.Sort(left)
	slices.Sort(right)
	fmt.Printf("List size: %d/%d\n", len(left), len(right))

	var summe int = 0
	for i := 0; i < len(left); i++ {
		summe += Abs(left[i] - right[i])
	}
	fmt.Printf("Result: %d\n", summe)
}
