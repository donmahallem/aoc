package day01

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/go/aoc_utils"
	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

func Part1(in io.Reader) (any, error) {
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
	slices.Sort(left)
	slices.Sort(right)

	var summe int = 0
	for i := range left {
		summe += int_util.AbsInt(left[i] - right[i])
	}
	return summe, nil
}
