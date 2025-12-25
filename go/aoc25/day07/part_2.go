package day07

import (
	"io"
)

func Part2(in io.Reader) (any, error) {
	inpData, err := parseInput(in)
	if err != nil {
		return 0, err
	}

	// If no splitter is hit: 1 timeline
	if inpData.startNode == nil {
		return 1, nil
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

	return dfs(inpData.startNode), nil
}
