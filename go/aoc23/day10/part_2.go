package day10

import (
	"io"
)

// collect the loop tiles and a lookup for nodes by position
func findLoop(start *Node) (map[point]bool, map[point]*Node) {
	loopSet := make(map[point]bool)
	loopNodes := make(map[point]*Node)
	if start == nil {
		return loopSet, loopNodes
	}
	// pick one neighbor to start walking the loop
	var cur *Node
	for _, nb := range []*Node{start.Up, start.Right, start.Down, start.Left} {
		if nb != nil {
			cur = nb
			break
		}
	}
	if cur == nil {
		return loopSet, loopNodes
	}
	prev := start
	loopSet[start.Position] = true
	loopNodes[start.Position] = start

	for cur != start {
		loopSet[cur.Position] = true
		loopNodes[cur.Position] = cur
		var next *Node
		for _, nb := range []*Node{cur.Up, cur.Right, cur.Down, cur.Left} {
			if nb != nil && nb != prev {
				next = nb
				break
			}
		}
		if next == nil {
			// not a loop; return what we have
			break
		}
		prev, cur = cur, next
	}
	return loopSet, loopNodes
}

// find number of enclosed cells in the loop using flood fill
func numberOfEnclosedCells(start *Node) int {
	loopSet, loopNodes := findLoop(start)
	if len(loopSet) == 0 {
		return 0
	}

	// bounds from loop tiles
	minX, maxX := start.Position.X, start.Position.X
	minY, maxY := start.Position.Y, start.Position.Y
	for p := range loopSet {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	w := maxX - minX + 1
	h := maxY - minY + 1

	// scaled grid: cells and edges; walls=true
	gw := w*2 + 1
	gh := h*2 + 1
	wall := make([][]bool, gh)
	seen := make([][]bool, gh)
	for y := 0; y < gh; y++ {
		wall[y] = make([]bool, gw)
		seen[y] = make([]bool, gw)
	}

	// mark loop centers and edges as walls
	for p, n := range loopNodes {
		cx := (p.X-minX)*2 + 1
		cy := (p.Y-minY)*2 + 1
		wall[cy][cx] = true
		for _, nb := range []*Node{n.Up, n.Right, n.Down, n.Left} {
			if nb == nil {
				continue
			}
			if !loopSet[nb.Position] {
				continue
			}
			dx := nb.Position.X - p.X // -1,0,1
			dy := nb.Position.Y - p.Y // -1,0,1
			wall[cy+dy][cx+dx] = true // the edge between centers
		}
	}

	// flood fill from outside (0,0) over non-wall
	type pt struct{ x, y int }
	q := make([]pt, 0, gw*gh/4)
	push := func(x, y int) {
		if x < 0 || y < 0 || x >= gw || y >= gh {
			return
		}
		if wall[y][x] || seen[y][x] {
			return
		}
		seen[y][x] = true
		q = append(q, pt{x, y})
	}
	push(0, 0)
	for i := 0; i < len(q); i++ {
		cur := q[i]
		push(cur.x+1, cur.y)
		push(cur.x-1, cur.y)
		push(cur.x, cur.y+1)
		push(cur.x, cur.y-1)
	}

	// count enclosed original cells (odd,odd) that are not wall and not seen
	enclosed := 0
	for y := range h {
		ey := y*2 + 1
		for x := range w {
			ex := x*2 + 1
			if !wall[ey][ex] && !seen[ey][ex] {
				enclosed++
			}
		}
	}
	return enclosed
}

func Part2(in io.Reader) int {
	start := ParseInput(in)
	return numberOfEnclosedCells(start)
}
