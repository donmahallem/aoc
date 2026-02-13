package day11

import (
	"bytes"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

func parseLine(in io.Reader) ([]int, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}
	splitData := bytes.Split(data, []byte(` `))
	retData := make([]int, len(splitData))
	for idx, val := range splitData {
		retData[idx] = 0
		for i := range len(val) {
			if val[i] < '0' || val[i] > '9' {
				return nil, aoc_utils.NewUnexpectedInputError(val[i])
			}
			retData[idx] = (10 * retData[idx]) + int(val[i]-'0')
		}
	}
	return retData, nil
}

type cacheKey struct {
	Stone int
	Depth uint8
}

func splitStone(stone int, depth uint8, cache map[cacheKey]int) int {
	if depth == 0 {
		return 1
	}
	cacheKey := cacheKey{Stone: stone, Depth: depth}
	if val, ok := cache[cacheKey]; ok {
		return val
	}
	var result int
	if stone == 0 {
		result = splitStone(1, depth-1, cache)
	} else if digits := int_util.CountDigits[int, int](stone); digits%2 == 0 {
		split := int_util.IntPow(10, digits/2)
		result = splitStone(stone/split, depth-1, cache) +
			splitStone(stone%split, depth-1, cache)
	} else {
		result = splitStone(stone*2024, depth-1, cache)
	}
	cache[cacheKey] = result
	return result
}

func splitStones(stones []int, depth uint8) int {
	cache := make(map[cacheKey]int, 512)
	result := 0
	for _, i := range stones {
		result += splitStone(i, depth, cache)
	}
	return result
}

func Part1(in io.Reader) (int, error) {
	data, err := parseLine(in)
	if err != nil {
		return 0, err
	}
	return splitStones(data, 25), nil
}
