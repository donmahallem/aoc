package day21

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
	"github.com/donmahallem/aoc/go/aoc_utils/bytes"
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
		if parsedInt, ok := bytes.ParseIntFromByte[uint](b); ok {
			val = (val * 10) + parsedInt
		}
	}
	return val
}

// Calculate amount of moves with the depth of directional keypads provided
// depth is inclusive(number of directional pads - 1)
func CalculateMoves(in io.Reader, depth uint8) uint {
	var total uint = 0
	cache := make(Cache)
	for pattern := range IterateInput(in) {
		total += ParseIntValue(&pattern) * WalkNumericSequence(&pattern, depth, &cache)
	}
	return total
}

func Part1(in io.Reader) uint {
	return CalculateMoves(in, 3)
}
