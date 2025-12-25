package day10

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type point = aoc_utils.Point[int]

type node struct {
	Position              point
	Mask                  uint8
	Right, Left, Up, Down *node
}

// Bitmask for the 4 directions
const (
	dirU = 1 << iota
	dirR
	dirD
	dirL
)

// Mapping from characters to direction bitmask
var char2mask = map[byte]uint8{
	'J': dirU | dirL,
	'-': dirR | dirL,
	'|': dirU | dirD,
	'L': dirU | dirR,
	'7': dirL | dirD,
	'F': dirR | dirD,
	'S': 0, // deduce later from neighbors
}

func parseInput(r io.Reader) (*node, error) {
	scanner := bufio.NewScanner(r)

	nodes := make(map[point]*node)
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
			n := &node{Position: p, Mask: m}
			nodes[p] = n
			masks[p] = m
		}
		y++
	}

	if !hasStart {
		return nil, io.EOF
	}

	// Deduce S mask from neighbors that actually connect back.
	start := nodes[startPos]
	dirOrder := []struct {
		dx, dy int
		bit    uint8
		opp    uint8
	}{
		{0, -1, dirU, dirD},
		{1, 0, dirR, dirL},
		{0, 1, dirD, dirU},
		{-1, 0, dirL, dirR},
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
		if (m & dirU) != 0 {
			up := point{X: p.X, Y: p.Y - 1}
			if nb, ok := nodes[up]; ok && (nb.Mask&dirD) != 0 {
				n.Up = nb
				nb.Down = n
			}
		}
		// Right
		if (m & dirR) != 0 {
			right := point{X: p.X + 1, Y: p.Y}
			if nb, ok := nodes[right]; ok && (nb.Mask&dirL) != 0 {
				n.Right = nb
				nb.Left = n
			}
		}
		// Down
		if (m & dirD) != 0 {
			down := point{X: p.X, Y: p.Y + 1}
			if nb, ok := nodes[down]; ok && (nb.Mask&dirU) != 0 {
				n.Down = nb
				nb.Up = n
			}
		}
		// Left
		if (m & dirL) != 0 {
			left := point{X: p.X - 1, Y: p.Y}
			if nb, ok := nodes[left]; ok && (nb.Mask&dirR) != 0 {
				n.Left = nb
				nb.Right = n
			}
		}
	}

	return start, nil
}
