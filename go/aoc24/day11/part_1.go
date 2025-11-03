package day11

import (
	"bytes"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils/math/log"
	"github.com/donmahallem/aoc/go/aoc_utils/math/pow"
)

func ParseLine(in io.Reader) ([]uint32, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}
	splitData := bytes.Split(data, []byte(` `))
	retData := make([]uint32, len(splitData))
	for idx, val := range splitData {
		retData[idx] = uint32(val[0] - '0')
		for i := 1; i < len(val); i++ {
			retData[idx] = (10 * retData[idx]) + uint32(val[i]-'0')
		}
	}
	return retData, nil
}

type CacheKey struct {
	Stone uint32
	Depth uint8
}

func SplitStone(stone uint32, depth uint8, cache map[CacheKey]int) int {
	if depth == 0 {
		return 1
	}
	cacheKey := CacheKey{Stone: stone, Depth: depth}
	if val, ok := cache[cacheKey]; ok {
		return val
	}
	var result int
	if stone == 0 {
		result = SplitStone(1, depth-1, cache)
	} else if digits := log.Log10Int(stone); digits%2 == 0 {
		split := pow.IntPow(10, digits/2)
		result = SplitStone(stone/split, depth-1, cache) +
			SplitStone(stone%split, depth-1, cache)
	} else {
		result = SplitStone(stone*2024, depth-1, cache)
	}
	cache[cacheKey] = result
	return result
}

func SplitStones(stones []uint32, depth uint8) int {
	cache := make(map[CacheKey]int, 512)
	result := 0
	for _, i := range stones {
		result += SplitStone(i, depth, cache)
	}
	return result
}

func Part1(in io.Reader) int {
	data, _ := ParseLine(in)

	return SplitStones(data, 25)
}
