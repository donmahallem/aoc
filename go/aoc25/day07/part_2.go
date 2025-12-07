package day07

import (
	"io"
)

func Part2(in io.Reader) int {
	startNode, _, _, _, _ := parseInput(in)

	// If no splitter is hit: 1 timeline
	if startNode == nil {
		return 1
	}

	cache := make(map[*node]int, 256)
	var dfs func(currentNode *node) int
	dfs = func(currentNode *node) int {
		if val, ok := cache[currentNode]; ok {
			return val
		}

		total := 0
		if currentNode.l != nil {
			total += dfs(currentNode.l)
		} else {
			// hit bottom
			total++
		}

		if currentNode.r != nil {
			total += dfs(currentNode.r)
		} else {
			// hit bottom
			total++
		}
		cache[currentNode] = total
		return total
	}

	return dfs(startNode)
}
