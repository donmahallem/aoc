package day09

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type inputRow []int
type input []inputRow

func parseInput(in io.Reader) (input, error) {
	s := bufio.NewScanner(in)
	result := make(input, 0)
	var err error
	for s.Scan() {
		parts := strings.Fields(s.Text())
		nums := make([]int, len(parts))
		for idx, item := range parts {
			nums[idx], err = strconv.Atoi(item)
			if err != nil {
				return nil, err
			}
		}
		result = append(result, nums)
	}
	return result, nil
}
