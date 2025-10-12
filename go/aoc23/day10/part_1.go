package day10

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

type point = aoc_utils.Point[int]

type Node struct {
	Position              point
	Mask                  uint8
	Right, Left, Up, Down *Node
}
type nodeMap = map[point]*Node

// Bitmask for the 4 directions
const (
	DirU = 1 << iota
	DirR
	DirD
	DirL
)

// Mapping from characters to direction bitmask
var char2mask = map[byte]uint8{
	'J': DirU | DirL,
	'-': DirR | DirL,
	'|': DirU | DirD,
	'L': DirU | DirR,
	'7': DirL | DirD,
	'F': DirR | DirD,
	'S': 0, // deduce later from neighbors
}

func ParseInput(r io.Reader) *Node {
	scanner := bufio.NewScanner(r)

	nodes := make(map[point]*Node)
	masks := make(map[point]uint8)
	var startPos point
	hasStart := false

	y := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		for x, ch := range line {
			if ch == '.' {
				continue
			}
			p := point{X: x, Y: y}
			m, ok := char2mask[ch]
			if !ok {
				continue
			}
			if ch == 'S' {
				startPos = p
				hasStart = true
			}
			n := &Node{Position: p, Mask: m}
			nodes[p] = n
			masks[p] = m
		}
		y++
	}

	if !hasStart {
		return nil
	}

	// Deduce S mask from neighbors that actually connect back.
	start := nodes[startPos]
	dirOrder := []struct {
		dx, dy int
		bit    uint8
		opp    uint8
	}{
		{0, -1, DirU, DirD},
		{1, 0, DirR, DirL},
		{0, 1, DirD, DirU},
		{-1, 0, DirL, DirR},
	}
	var sMask uint8
	for _, d := range dirOrder {
		np := point{X: startPos.X + d.dx, Y: startPos.Y + d.dy}
		if nm, ok := masks[np]; ok && (nm&d.opp) != 0 {
			sMask |= d.bit
		}
	}
	start.Mask = sMask
	masks[startPos] = sMask

	// Connect nodes only if both sides agree.
	for p, n := range nodes {
		m := n.Mask
		// Up
		if (m & DirU) != 0 {
			up := point{X: p.X, Y: p.Y - 1}
			if nb, ok := nodes[up]; ok && (nb.Mask&DirD) != 0 {
				n.Up = nb
				nb.Down = n
			}
		}
		// Right
		if (m & DirR) != 0 {
			right := point{X: p.X + 1, Y: p.Y}
			if nb, ok := nodes[right]; ok && (nb.Mask&DirL) != 0 {
				n.Right = nb
				nb.Left = n
			}
		}
		// Down
		if (m & DirD) != 0 {
			down := point{X: p.X, Y: p.Y + 1}
			if nb, ok := nodes[down]; ok && (nb.Mask&DirU) != 0 {
				n.Down = nb
				nb.Up = n
			}
		}
		// Left
		if (m & DirL) != 0 {
			left := point{X: p.X - 1, Y: p.Y}
			if nb, ok := nodes[left]; ok && (nb.Mask&DirR) != 0 {
				n.Left = nb
				nb.Right = n
			}
		}
	}

	return start
}

func loopLengthFrom(start *Node) int {
	if start == nil {
		return 0
	}
	// Pick one of the two neighbors of S
	var cur *Node
	for _, nb := range []*Node{start.Up, start.Right, start.Down, start.Left} {
		if nb != nil {
			cur = nb
			break
		}
	}
	if cur == nil {
		return 0
	}
	prev := start
	length := 1
	for cur != start {
		var next *Node
		for _, nb := range []*Node{cur.Up, cur.Right, cur.Down, cur.Left} {
			if nb != nil && nb != prev {
				next = nb
				break
			}
		}
		if next == nil {
			return 0 // not a loop
		}
		prev, cur = cur, next
		length++
	}
	return length
}

func Part1(in io.Reader) int {
	start := ParseInput(in)
	loopSize := loopLengthFrom(start)
	return loopSize / 2
}
