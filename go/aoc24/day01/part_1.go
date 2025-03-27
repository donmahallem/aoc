package day01

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/aoc_utils"
)

func Part1(in io.Reader) int {
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
	slices.Sort(left)
	slices.Sort(right)
	fmt.Printf("List size: %d/%d\n", len(left), len(right))

	var summe int = 0
	for i := 0; i < len(left); i++ {
		summe += aoc_utils.Abs(left[i] - right[i])
	}
	return summe
}
