package day02

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

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

func Part2(in *os.File) {
	s := bufio.NewScanner(in)
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
