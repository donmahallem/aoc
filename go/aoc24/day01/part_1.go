package day01

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/aoc_utils/math/abs"
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

	var summe int = 0
	for i := range len(left) {
		summe += abs.AbsInt(left[i] - right[i])
	}
	return summe
}
