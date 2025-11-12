package day23

import "fmt"

func findStart(parseData []cell, w int) int {
	for idx := 0; idx < w; idx++ {
		if parseData[idx]&dirExitBottom != 0 {
			return idx
		}
	}
	return -1
}

func findEnd(parseData []cell, w int) int {
	rowStart := len(parseData) - w
	for idx := len(parseData) - 1; idx >= rowStart; idx-- {
		if parseData[idx]&dirEntryTop != 0 {
			return idx
		}
	}
	return -1
}

type edge struct {
	distance int
	toIdx    int
}

type node struct {
	idx       int
	neighbors []edge
}

func createGraph(parseData []cell, w, h int) (map[int]*node, int, int) {
	if len(parseData) == 0 || w <= 0 || h <= 0 {
		return map[int]*node{}, -1, -1
	}

	startIdx := findStart(parseData, w)
	if startIdx == -1 {
		return map[int]*node{}, -1, -1
	}
	endIdx := findEnd(parseData, w)
	if endIdx == -1 {
		return map[int]*node{}, -1, -1
	}

	graph := make(map[int]*node)
	graph[startIdx] = &node{idx: startIdx}
	queue := []int{startIdx}
	visited := make([]bool, len(parseData))

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if visited[current] {
			continue
		}
		visited[current] = true

		currentNode := graph[current]
		if currentNode == nil {
			currentNode = &node{idx: current}
			graph[current] = currentNode
		}

		exitDirections := parseData[current] & dirExitAll
		for remainingExitDirections := exitDirections; remainingExitDirections != 0; {
			dirMask := lowestBit(remainingExitDirections)
			remainingExitDirections &^= dirMask

			_, ok := stepIndex(current, dirMask, w, h)
			if !ok {
				continue
			}

			destIdx, distance := walkCorridor(current, dirMask, parseData, w, h, startIdx, endIdx)
			if destIdx == -1 || destIdx == current || distance == 0 {
				continue
			}

			//if !edgeExists(currentNode.neighbors, destIdx) {
			//	currentNode.neighbors = append(currentNode.neighbors, edge{distance: distance, toIdx: destIdx})
			//
			upsertEdge(&currentNode.neighbors, destIdx, distance)

			if _, found := graph[destIdx]; !found {
				graph[destIdx] = &node{idx: destIdx}
			}
			if !visited[destIdx] {
				queue = append(queue, destIdx)
			}
		}
	}

	return graph, startIdx, endIdx
}

func upsertEdge(edges *[]edge, to int, distance int) {
	for i, e := range *edges {
		if e.toIdx == to {
			if distance > e.distance {
				fmt.Printf("found\n")
				(*edges)[i].distance = distance
			}
			return
		}
	}
	*edges = append(*edges, edge{toIdx: to, distance: distance})
}

func edgeExists(edges []edge, to int) bool {
	for _, e := range edges {
		if e.toIdx == to {
			return true
		}
	}
	return false
}

// returns the lowest set bit in the mask
func lowestBit(mask cell) cell {
	return mask & -mask
}

// returns the index stepped in direction dirMask and whether the step is valid
func stepIndex(idx int, dirMask cell, w, h int) (int, bool) {
	switch dirMask {
	case dirExitTop:
		if idx < w {
			return -1, false
		}
		return idx - w, true
	case dirExitBottom:
		if idx/w >= h-1 {
			return -1, false
		}
		return idx + w, true
	case dirExitLeft:
		if idx%w == 0 {
			return -1, false
		}
		return idx - 1, true
	case dirExitRight:
		if idx%w == w-1 {
			return -1, false
		}
		return idx + 1, true
	default:
		return -1, false
	}
}

// returns the opposite direction mask
func opposite(dirMask cell) cell {
	switch dirMask {
	case dirExitTop:
		return dirExitBottom
	case dirExitBottom:
		return dirExitTop
	case dirExitLeft:
		return dirExitRight
	case dirExitRight:
		return dirExitLeft
	default:
		return 0
	}
}

// walks the corridor starting at startIdx in direction dirMask
// returns the next node index and the distance walked
func walkCorridor(startIdx int, dirMask cell, parseData []cell, w, h int, startNode, endNode int) (int, int) {
	entryMask := entryMaskFor(dirMask)
	nextIdx, ok := stepIndex(startIdx, dirMask, w, h)
	if !ok || entryMask == 0 || parseData[nextIdx] == dirMaskNone || parseData[nextIdx]&entryMask == 0 {
		return -1, 0
	}

	distance := 1
	incoming := opposite(dirMask)
	limit := len(parseData)

	for steps := 0; steps < limit; steps++ {
		mask := parseData[nextIdx]

		if isNode(nextIdx, mask, incoming, startNode, endNode) {
			return nextIdx, distance
		}

		exits := (mask & dirExitAll) &^ incoming
		if exits == 0 {
			return -1, 0
		}

		nextDir := lowestBit(exits)
		entryMask = entryMaskFor(nextDir)
		targetIdx, ok := stepIndex(nextIdx, nextDir, w, h)
		if !ok || entryMask == 0 || parseData[targetIdx] == dirMaskNone || parseData[targetIdx]&entryMask == 0 {
			return -1, 0
		}

		incoming = opposite(nextDir)
		nextIdx = targetIdx
		distance++
	}

	return -1, 0
}

func isNode(idx int, mask cell, incoming cell, startIdx, endIdx int) bool {
	if idx == startIdx || idx == endIdx {
		return true
	}
	exits := (mask & dirExitAll) &^ incoming
	return bitCount(exits) != 1
}

func entryMaskFor(dirMask cell) cell {
	switch dirMask {
	case dirExitTop:
		return dirEntryBottom
	case dirExitBottom:
		return dirEntryTop
	case dirExitLeft:
		return dirEntryRight
	case dirExitRight:
		return dirEntryLeft
	default:
		return 0
	}
}

// counts the number of set bits in the mask
func bitCount(mask cell) int {
	count := 0
	for mask != 0 {
		mask &= mask - 1
		count++
	}
	return count
}
