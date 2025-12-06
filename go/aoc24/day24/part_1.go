package day24

import (
	"io"
)

func translateId(id uint32) uint8 {
	// interpret the two low bytes as ASCII digits "xy" -> x*10 + y
	low := byte(id & 0xFF)        // third char
	mid := byte((id >> 8) & 0xFF) // second char
	d0 := low - '0'
	d1 := mid - '0'
	return d1*10 + d0
}

func Part1(in io.Reader) uint64 {
	data := parseInput(in)

	output := uint64(0)
	for id, node := range data.nodes {
		if id>>16 == uint32('z') {
			if node.Evaluate() {
				output += uint64(1) << uint(translateId(id))
			}
		}
	}
	return output
}
