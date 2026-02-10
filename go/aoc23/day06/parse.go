package day06

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/go/aoc_utils/bytes"
)

type race struct {
	Time     int
	Distance int
}

func parseInput(in io.Reader) []race {
	row1 := make([]int, 0, 4)
	row2 := make([]int, 0, 4)

	s := bufio.NewScanner(in)
	// read first row
	if !s.Scan() {
		return nil
	}
	for idx, item := range strings.Fields(s.Text()) {
		if idx < 1 {
			continue
		}
		val, _ := strconv.Atoi(item)
		row1 = append(row1, val)
	}
	// read second row
	if !s.Scan() {
		return nil
	}
	for idx, item := range strings.Fields(s.Text()) {
		if idx < 1 {
			continue
		}
		val, _ := strconv.Atoi(item)
		row2 = append(row2, val)
	}
	// ensure matching lengths
	if len(row1) == 0 || len(row1) != len(row2) {
		return nil
	}
	races := make([]race, len(row1))
	for idx := range row1 {
		races[idx] = race{Time: row1[idx], Distance: row2[idx]}
	}
	return races
}

func parseInputPart2(in io.Reader) race {
	var time int
	s := bufio.NewScanner(in)
	firstNumber := true
	currentNumber := 0
	for s.Scan() {
		b := s.Bytes()
		for _, c := range b {
			if parsedInt, ok := bytes.ParseIntFromByte[int](c); ok {
				currentNumber *= 10
				currentNumber += parsedInt
			}
		}
		if firstNumber {
			time = currentNumber
			currentNumber = 0
			firstNumber = false
		} else {
			break
		}
	}
	return race{Time: time, Distance: currentNumber}
}
