package day11

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/donmahallem/aoc/utils"
)

func ParseLine(in io.Reader) ([]int, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}
	splitData := bytes.Split(data, []byte(` `))
	retData := make([]int, len(splitData))
	for idx, val := range splitData {
		retData[idx] = int(val[0] - '0')
		for i := 1; i < len(val); i++ {
			retData[idx] = (10 * retData[idx]) + int(val[i]-'0')
		}
	}
	return retData, nil
}

func SplitStone(stone int, depth int, cache *map[[2]int]int) int {
	if depth == 0 {
		return 1
	}
	cacheKey := [2]int{stone, depth}
	if val, ok := (*cache)[cacheKey]; ok {
		return val
	}
	var result int
	if stone == 0 {
		result = SplitStone(1, depth-1, cache)
	} else if digits := utils.Log10Int(stone); digits%2 == 0 {
		split := utils.IntPow(10, digits/2)
		result = SplitStone(stone/split, depth-1, cache) +
			SplitStone(stone%split, depth-1, cache)
	} else {
		result = SplitStone(stone*2024, depth-1, cache)
	}
	(*cache)[cacheKey] = result
	return result
}

func SplitStones(stones []int, depth int) int {
	cache := make(map[[2]int]int)
	result := 0
	for _, i := range stones {
		result += SplitStone(i, depth, &cache)
	}
	return result
}

func Part1(in *os.File) {
	data, _ := ParseLine(in)

	fmt.Printf("%d\n", SplitStones(data, 25))
}
