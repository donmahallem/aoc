package day08

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type pathInstructions []bool

type node struct {
	Left, Right *node
	Id          uint32
	EndsInZ     bool
	EndsInA     bool
}
type input struct {
	// true if right else false
	Instructions pathInstructions
	Start        *node
	End          *node
	Nodes        map[uint32]*node
}

func charVal(b byte) (uint32, error) {
	if b >= '0' && b <= '9' {
		return uint32(b - '0'), nil
	}
	if b >= 'A' && b <= 'Z' {
		return uint32(b-'A') + 10, nil
	}
	return 0, aoc_utils.NewUnexpectedInputError(b)
}

func hashNodeId(id []byte) uint32 {
	// base 36 encoding 3 chars -> val in [0, 36^3)
	v0, _ := charVal(id[0])
	v1, _ := charVal(id[1])
	v2, _ := charVal(id[2])
	return (v0 * 36 * 36) + (v1 * 36) + v2
}

func getOrCreateNode(nodeMap *map[uint32]*node, strId []byte) (*node, error) {
	// expect exactly 3 bytes for an ID
	if len(strId) < 3 {
		return nil, aoc_utils.NewParseError("Node ID must be at least 3 characters long", nil)
	}
	// validate and compute id
	for i := 0; i < 3; i++ {
		ch := strId[i]
		if _, err := charVal(ch); err != nil {
			return nil, err
		}
	}
	id := hashNodeId(strId[:3])
	if n, ok := (*nodeMap)[id]; ok {
		return n, nil
	}
	n := &node{Id: id}
	n.EndsInZ = strId[2] == 'Z'
	n.EndsInA = strId[2] == 'A'
	(*nodeMap)[id] = n
	return n, nil
}

func parseId(b []byte, offset int) (uint32, error) {
	if len(b) < offset+3 {
		return 0, aoc_utils.NewParseError("ID must be at least 3 characters long", nil)
	}
	id := uint32(0)
	for i := 0; i < 3; i++ {
		ch := b[offset+i]
		val, err := charVal(ch)
		if err != nil {
			return 0, err
		}
		id = id*36 + val
	}
	return id, nil
}

func parseInstructions(line []byte) (pathInstructions, error) {
	instructions := make(pathInstructions, 0, len(line))
	for _, ch := range line {
		if ch == 'R' {
			instructions = append(instructions, true)
		} else if ch == 'L' {
			instructions = append(instructions, false)
		} else {
			return nil, aoc_utils.NewUnexpectedInputError(ch)
		}
	}
	return instructions, nil
}

func parsePairLine(line []byte) (curr, left, right []byte, err error) {
	n := len(line)
	i := 0
	// helper: skip spaces
	skipSpaces := func() {
		for i < n && line[i] == ' ' {
			i++
		}
	}
	// parse a 3-char ID at current index
	parseIDAt := func() ([]byte, error) {
		if i+3 > n {
			return nil, aoc_utils.NewParseError("malformed node line", nil)
		}
		for k := 0; k < 3; k++ {
			if _, err := charVal(line[i+k]); err != nil {
				return nil, err
			}
		}
		id := line[i : i+3]
		i += 3
		return id, nil
	}

	// start: curr id
	skipSpaces()
	currId, err := parseIDAt()
	if err != nil {
		return nil, nil, nil, err
	}
	skipSpaces()
	// expect '='
	if i >= n || line[i] != '=' {
		return nil, nil, nil, aoc_utils.NewParseError("malformed node line: missing '='", nil)
	}
	i++
	skipSpaces()
	// expect '('
	if i >= n || line[i] != '(' {
		return nil, nil, nil, aoc_utils.NewParseError("malformed node line: missing '('", nil)
	}
	i++
	skipSpaces()
	leftId, err := parseIDAt()
	if err != nil {
		return nil, nil, nil, err
	}
	skipSpaces()
	// expect ','
	if i >= n || line[i] != ',' {
		return nil, nil, nil, aoc_utils.NewParseError("malformed node line: missing ','", nil)
	}
	i++
	skipSpaces()
	rightId, err := parseIDAt()
	if err != nil {
		return nil, nil, nil, err
	}
	skipSpaces()
	// expect ')'
	if i >= n || line[i] != ')' {
		return nil, nil, nil, aoc_utils.NewParseError("malformed node line: missing ')'", nil)
	}
	i++
	skipSpaces()
	if i != n {
		// trailing garbage
		return nil, nil, nil, aoc_utils.NewParseError("malformed node line: trailing characters", nil)
	}
	return currId, leftId, rightId, nil
}

func parseInput(in io.Reader) (*input, error) {
	startId := hashNodeId([]byte{'A', 'A', 'A'})
	endId := hashNodeId([]byte{'Z', 'Z', 'Z'})
	s := bufio.NewScanner(in)
	inp := input{}
	inp.Nodes = make(map[uint32]*node, 64)
	pNodeMap := &inp.Nodes
	instructionsRead := false
	var instructionsParsed pathInstructions
	var err error
	for s.Scan() {
		line := s.Bytes()
		if len(line) == 0 {
			continue
		} else if !instructionsRead {
			instructionsParsed, err = parseInstructions(line)
			if err != nil {
				return nil, err
			}
			instructionsRead = true
			continue
		} else {
			currBytes, leftBytes, rightBytes, err := parsePairLine(line)
			if err != nil {
				return nil, err
			}
			currentNode, err := getOrCreateNode(pNodeMap, currBytes)
			if err != nil {
				return nil, err
			}
			leftNode, err := getOrCreateNode(pNodeMap, leftBytes)
			if err != nil {
				return nil, err
			}
			rightNode, err := getOrCreateNode(pNodeMap, rightBytes)
			if err != nil {
				return nil, err
			}
			currentNode.Left = leftNode
			currentNode.Right = rightNode
			if currentNode.Id == startId {
				inp.Start = currentNode
			} else if currentNode.Id == endId {
				inp.End = currentNode
			}
		}

	}
	return &input{
		Instructions: instructionsParsed,
		Nodes:        inp.Nodes,
		Start:        inp.Start,
		End:          inp.End,
	}, nil
}
