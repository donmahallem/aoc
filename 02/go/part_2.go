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

func checkRecursive(l []int, dir int) {

}

func checkLine(l []int) bool {
	var dir bool = false
	for i := 1; i < len(l); i++ {
		var diff int = l[i] - l[i-1]
		if diff == 0 || Abs(diff) > 3 {
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

func checkVariations(l []int) bool {
	if checkLine(l) {
		return true
	}
	for i := 0; i < len(l); i++ {
		var testSlice = slices.Concat(l[0:i], l[i+1:])
		if checkLine((testSlice)) {
			return true
		}
	}
	return false
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var goodLines = 0
	var totalLines = 0
	for s.Scan() {
		var line = strings.Split(s.Text(), " ")
		var parsedLine = make([]int, len(line))
		for idx, item := range line {
			var val, _ = strconv.Atoi(item)
			parsedLine[idx] = val
		}
		if checkVariations((parsedLine)) {
			goodLines++
		}
		totalLines++
	}
	fmt.Printf("List size: %d/%d\n", goodLines, totalLines)
}
