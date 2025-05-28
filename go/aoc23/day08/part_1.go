package day08

import (
	"bufio"
	"io"
)

type PathInstructions []bool

type Node struct {
	Left, Right *Node
	Id          uint32
	EndsInZ     bool
	EndsInA     bool
}
type Input struct {
	// true if right else false
	Instructions PathInstructions
	Start        *Node
	End          *Node
	Nodes        map[uint32]*Node
}

func hashNodeId(id []byte) uint32 {
	return uint32(id[0]-'A')<<12 + uint32(id[1]-'A')<<6 + uint32(id[2]-'A')
}

func getOrCreateNode(nodeMap *map[uint32]*Node, strId []byte) *Node {
	id := hashNodeId(strId)
	if node, ok := (*nodeMap)[id]; ok {
		return node
	}
	node := &Node{}
	node.Id = id
	node.EndsInZ = strId[2] == 'Z'
	node.EndsInA = strId[2] == 'A'
	(*nodeMap)[id] = node
	return node
}

func ParseInput(in io.Reader) Input {
	startId := hashNodeId([]byte{'A', 'A', 'A'})
	endId := hashNodeId([]byte{'Z', 'Z', 'Z'})
	s := bufio.NewScanner(in)
	inp := Input{}
	inp.Nodes = make(map[uint32]*Node, 64)
	pNodeMap := &inp.Nodes
	instructionsRead := false
	for s.Scan() {
		line := s.Bytes()
		if len(line) == 0 {
			continue
		} else if !instructionsRead {
			inp.Instructions = make(PathInstructions, len(line))
			for idx, ch := range line {
				if ch == 'R' {
					inp.Instructions[idx] = true
				}
			}
			instructionsRead = true
		} else {
			currentNode := getOrCreateNode(pNodeMap, line)
			currentNode.Left = getOrCreateNode(pNodeMap, line[7:10])
			currentNode.Right = getOrCreateNode(pNodeMap, line[12:15])
			if currentNode.Id == startId {
				inp.Start = currentNode
			} else if currentNode.Id == endId {
				inp.End = currentNode
			}
		}
	}
	return inp
}

func Part1(in io.Reader) int {
	games := ParseInput(in)
	steps := 0
	numInstructions := len(games.Instructions)
	current := games.Start
	for ; ; steps++ {
		if current == games.End {
			break
		}
		if games.Instructions[steps%numInstructions] {
			current = current.Right
		} else {
			current = current.Left
		}
	}
	return steps
}
