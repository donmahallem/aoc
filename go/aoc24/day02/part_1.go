package day02

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

func checkLine(l []int) bool {
	var dir bool = false
	for i := 1; i < len(l); i++ {
		var diff int = l[i] - l[i-1]
		if diff == 0 || int_util.AbsInt(diff) > 3 {
			return false
		}
		if i > 1 {
			if dir && diff < 0 {
				return false
			} else if !dir && diff > 0 {
				return false
			}
		} else {
			if diff < 0 {
				dir = false
			} else {
				dir = true
			}
		}
	}
	return true
}

func Part1(in io.Reader) (any, error) {
	s := bufio.NewScanner(in)
	var goodLines = 0
	for s.Scan() {
		var line = strings.Split(s.Text(), " ")
		var parsedLine = make([]int, len(line))
		for idx, item := range line {
			var val, _ = strconv.Atoi(item)
			parsedLine[idx] = val
		}
		if checkLine((parsedLine)) {
			goodLines++
		}
	}
	return goodLines, nil
}
