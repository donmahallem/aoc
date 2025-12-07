package day07

import (
	"io"
)

func Part2(in io.Reader) int {
	splitterMap, startX, startY, width, height := parseInput(in)

	cache := make(map[int]int, len(splitterMap))
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if y == height-1 {
			return 1
		}
		pos := y*width + x
		if _, ok := splitterMap[pos]; ok {
			if val, ok := cache[pos]; ok {
				return val
			}
			total := 0
			if x > 0 {
				total += dfs(x-1, y+1)
			}
			if x < width-1 {
				total += dfs(x+1, y+1)
			}
			cache[pos] = total
			return total
		}
		return dfs(x, y+1)
	}

	totalCount := dfs(startX, startY)

	return totalCount
}
