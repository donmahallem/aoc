package day07

import (
	"io"
)

func Part1(in io.Reader) int {
	splitterMap, _, _, width, height := parseInput(in)

	rays := make([]*node, 128)
	head, tail := 0, 1
	rays[0] = splitterMap

	visited := make(map[*node]bool, width*height/4)

	for head < tail {
		node := rays[head]
		head++

		for {
			if visited[node] {
				break
			}
			visited[node] = true

			if node.l == nil && node.r == nil {
				break
			}

			// branch down-left
			if node.l != nil {
				if tail == len(rays) {
					rays = append(rays, node.l)
				} else {
					rays[tail] = node.l
				}
				tail++
			}
			// branch down-right
			if node.r != nil {
				if tail == len(rays) {
					rays = append(rays, node.r)
				} else {
					rays[tail] = node.r
				}
				tail++
			}
			break
		}
	}

	return len(visited)

}
