package day08

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

func GetCycleSize(instructions PathInstructions, node *Node) uint {
	current := node
	for i := 0; ; i++ {
		if i > 0 && current.EndsInZ && i%len(instructions) == 0 {
			return uint(i)
		}
		if instructions[i%len(instructions)] {
			current = current.Right
		} else {
			current = current.Left
		}
	}
}

func Part2(in io.Reader) uint {
	games := ParseInput(in)
	var currentBase uint = 1
	for _, node := range games.Nodes {
		if node.EndsInA {
			currentBase = int_util.LcmInt(currentBase, GetCycleSize(games.Instructions, node))
		}
	}
	return currentBase
}
