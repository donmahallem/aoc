package day24

import (
	"bufio"
	_ "embed"
	"io"
)

type nodeId = uint32

func HashId(a []byte) nodeId {
	return nodeId(a[0])<<16 | nodeId(a[1])<<8 | nodeId(a[2])
}
func HashIdA(a, b, c byte) nodeId {
	return nodeId(a)<<16 | nodeId(b)<<8 | nodeId(c)
}

func UnhashId(id nodeId) [3]byte {
	return [3]byte{byte(id >> 16), byte((id >> 8) & 0xFF), byte(id & 0xFF)}
}

type operation struct {
	nodeA, nodeB  nodeId
	operationType op
}

type node struct {
	id nodeId
}

type operationNode struct {
	operation operation
}

type importData struct {
	inputNodes  map[nodeId]bool
	connections map[nodeId]map[nodeId]struct{}
	nodes       map[nodeId]evalNode
	output      nodeId
}

func parseInput(in io.Reader) importData {
	s := bufio.NewScanner(in)
	initialBlock := true
	data := importData{
		inputNodes:  make(map[nodeId]bool),
		connections: make(map[nodeId]map[nodeId]struct{}),
		nodes:       make(map[nodeId]evalNode),
	}
	for s.Scan() {
		line := s.Bytes()
		if len(line) == 0 {
			initialBlock = false
			continue
		}
		if initialBlock {
			// first block is input nodes
			key := HashId(line[0:3])
			data.inputNodes[key] = line[5] == '1'
			data.nodes[key] = &valueNode{value: line[5] == '1'}

		} else {
			// second block is connections
			currentOperation := operation{
				nodeA:         HashId(line[0:3]),
				nodeB:         21,
				operationType: OPERATION_ADD,
			}
			offset := 4
			switch line[4] {
			case 'A':
				currentOperation.operationType = OPERATION_ADD
				offset += 4
			case 'O':
				currentOperation.operationType = OPERATION_OR
				offset += 3
			case 'X':
				currentOperation.operationType = OPERATION_XOR
				offset += 4
			}
			currentOperation.nodeB = HashId(line[offset : offset+3])

			outputNodeId := HashId(line[offset+7 : offset+10])
			ensureNode := func(id nodeId) evalNode {
				if n, ok := data.nodes[id]; ok && n != nil {
					return n
				}
				v := &nilableNode{}
				data.nodes[id] = v
				return v
			}

			left := ensureNode(currentOperation.nodeA)
			right := ensureNode(currentOperation.nodeB)

			if cur, ok := data.nodes[outputNodeId]; ok {
				// already exists, should be nilable node
				nilable := cur.(*nilableNode)
				nilable.left = left
				nilable.right = right
				nilable.opType = currentOperation.operationType

			} else {
				data.nodes[outputNodeId] = &nilableNode{
					left:   left,
					right:  right,
					opType: currentOperation.operationType,
				}
			}
		}
	}
	return data
}
