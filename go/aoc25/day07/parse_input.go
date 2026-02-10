package day07

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type node struct {
	l, r *node
	x    int
}

type inputData struct {
	startNode *node
	startX    int
	startY    int
	width     int
	height    int
}

func parseInput(in io.Reader) (*inputData, error) {
	s := bufio.NewScanner(in)
	nodes := make([]node, 0, 1024)
	var startX, startY, height int = -1, -1, 0
	width := -1
	var startNode *node

	for ; s.Scan(); height++ {
		line := s.Bytes()
		if width < 0 {
			width = len(line)
		} else if width != len(line) {
			return nil, aoc_utils.NewParseError("inconsistent line widths", nil)
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

	if startX < 0 {
		return nil, aoc_utils.NewParseError("missing start node", nil)
	}
	if width <= 0 || height <= 0 {
		return nil, aoc_utils.NewParseError("invalid dimensions", nil)
	}

	nodeCache := make([]*node, width)

	if len(nodes) == 0 {
		return nil, aoc_utils.NewParseError("no nodes found", nil)
	}
	// iterate backwards to link nodes
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

	if startNode == nil {
		return nil, aoc_utils.NewParseError("start node not found or reached", nil)
	}
	return &inputData{
		startNode: startNode,
		startX:    startX,
		startY:    startY,
		width:     width,
		height:    height,
	}, nil
}
