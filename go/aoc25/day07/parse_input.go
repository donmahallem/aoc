package day07

import (
	"bufio"
	"io"
)

type splitterPositions map[int]struct{}

type node struct {
	l, r *node
	x    int
}

func parseInput(in io.Reader) (*node, int, int, int, int) {
	s := bufio.NewScanner(in)
	nodes := make([]node, 0, 1024)
	var startX, startY, height int
	width := -1
	var startNode *node

	for ; s.Scan(); height++ {
		line := s.Bytes()
		if width < 0 {
			width = len(line)
		}
		for x, c := range line {
			switch c {
			case '^':
				nodes = append(nodes, node{x: x})
			case 'S':
				startX, startY = x, height
			}
		}
	}

	nodeCache := make([]*node, width)

	for i := len(nodes) - 1; i >= 0; i-- {
		n := &nodes[i]
		x := n.x
		if x == startX {
			startNode = n
		}
		if x > 0 {
			n.l = nodeCache[x-1]
		}
		if x < width-1 {
			n.r = nodeCache[x+1]
		}
		nodeCache[x] = n
	}

	return startNode, startX, startY, width, height
}
