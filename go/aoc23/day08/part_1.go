package day08

import (
	"bufio"
	"io"
)

type PathInstructions []bool

type Node struct {
	Left, Right *Node
}
type Input struct {
	Instruktions PathInstructions
	Start        *Node
	End          *Node
}

func hashNodeId(id []byte) uint32 {
	return uint32(id[0]-'A')<<12 + uint32(id[1]-'A')<<6 + uint32(id[2]-'A')
}

func getOrCreateNode(nodeMap *map[uint32]*Node, id uint32) *Node {
	if node, ok := (*nodeMap)[id]; ok {
		return node
	}
	node := &Node{}
	(*nodeMap)[id] = node
	return node
}

func ParseInput(in io.Reader) Input {
	startId := hashNodeId([]byte{'A', 'A', 'A'})
	endId := hashNodeId([]byte{'Z', 'Z', 'Z'})
	s := bufio.NewScanner(in)
	inp := Input{}
	nodeMap := make(map[uint32]*Node, 64)
	pNodeMap := &nodeMap
	instructionsRead := false
	for s.Scan() {
		line := s.Bytes()
		if len(line) == 0 {
			continue
		} else if !instructionsRead {
			inp.Instruktions = make(PathInstructions, len(line))
			for idx, ch := range line {
				if ch == 'R' {
					inp.Instruktions[idx] = true
				}
			}
			instructionsRead = true
		} else {
			currentId := hashNodeId(line)
			leftId := hashNodeId(line[7:10])
			rightId := hashNodeId(line[12:15])
			currentNode := getOrCreateNode(pNodeMap, currentId)
			currentNode.Left = getOrCreateNode(pNodeMap, leftId)
			currentNode.Right = getOrCreateNode(pNodeMap, rightId)
			if currentId == startId {
				inp.Start = currentNode
			} else if currentId == endId {
				inp.End = currentNode
			}
		}
	}
	return inp
}

func Part1(in io.Reader) int {
	games := ParseInput(in)
	steps := 0
	numInstructions := len(games.Instruktions)
	current := games.Start
	for ; ; steps++ {
		if current == games.End {
			break
		}
		if games.Instruktions[steps%numInstructions] {
			current = current.Right
		} else {
			current = current.Left
		}
	}
	return steps
}
