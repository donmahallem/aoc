package day11

import (
	"io"
)

func Part1(in io.Reader) uint64 {
	inputMap := parseInput(in)

	targetKey := uint64('o')<<16 + uint64('u')<<8 + uint64('t')
	visited := make(map[uint64]uint64, 1024)
	var dfs func(node uint64) uint64
	dfs = func(node uint64) uint64 {
		if node == targetKey {
			return 1
		}
		if val, ok := visited[node]; ok {
			return val
		}
		currentValue := uint64(0)
		for _, target := range inputMap[node] {
			currentValue += dfs(target)
		}
		visited[node] = currentValue // Store the result in the cache
		return currentValue
	}
	return dfs(uint64('y')<<16 + uint64('o')<<8 + uint64('u'))
}
