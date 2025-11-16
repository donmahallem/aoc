package day23

func dfsIterative(graph map[int]*node, start int, target int) int {
	startNode := graph[start]
	if startNode == nil {
		return 0
	}
	if start == target {
		return 0
	}

	maxIdx := 0
	for k := range graph {
		if k > maxIdx {
			maxIdx = k
		}
	}

	visited := make([]bool, maxIdx+1)

	type frame struct {
		idx      int
		edgeIdx  int
		distance int
	}

	stack := make([]frame, 0, len(graph))
	stack = append(stack, frame{idx: start, edgeIdx: 0, distance: 0})
	visited[start] = true
	longest := 0

	for len(stack) > 0 {
		top := &stack[len(stack)-1]
		currentNode := graph[top.idx]

		if top.edgeIdx >= len(currentNode.neighbors) {
			visited[top.idx] = false // Backtrack
			stack = stack[:len(stack)-1]
			continue
		}

		edge := currentNode.neighbors[top.edgeIdx]
		top.edgeIdx++

		if visited[edge.toIdx] {
			continue
		}

		newDistance := top.distance + edge.distance
		if edge.toIdx == target {
			if newDistance > longest {
				longest = newDistance
			}
			continue
		}

		visited[edge.toIdx] = true
		stack = append(stack, frame{idx: edge.toIdx, edgeIdx: 0, distance: newDistance})
	}

	return longest
}
