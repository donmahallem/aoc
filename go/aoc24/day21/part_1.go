package day21

import (
	"bufio"
	"fmt"
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

type Point = aoc_utils.Point[int8]

func IterateInput(in io.Reader) chan []byte {
	c := make(chan []byte)
	go func() {
		s := bufio.NewScanner(in)
		for s.Scan() {
			c <- s.Bytes()
		}
		close(c)
	}()
	return c
}

func ParseIntValue(data *[]byte) uint {
	var val uint = 0
	for _, b := range *data {
		if b >= '0' && b <= '9' {
			val = (val * 10) + uint(b-'0')
		}
	}
	return val
}

func CalculateMoves(in io.Reader, depth uint8) uint {
	var total uint = 0
	for pattern := range IterateInput(in) {
		total += ParseIntValue(&pattern) * WalkNumericSequence(&pattern, depth)
	}
	fmt.Printf("Cache size: %d\n", len(cache))
	return total
}

func Part1(in io.Reader) uint {
	return CalculateMoves(in, 3)
}
