package day07

import (
	"io"
)

func Part1(in io.Reader) (any, error) {
	data, err := parseInput(in)
	if err != nil {
		return 0, err
	}

	rays := make([]*node, 128)
	head, tail := 0, 1
	rays[0] = data.startNode

	visited := make(map[*node]bool, data.width*data.height/4)

	for head < tail {
		n := rays[head]
		head++

		for {
			if visited[n] {
				break
			}
			visited[n] = true

			if n.l == nil && n.r == nil {
				break
			}

			// branch down-left
			if n.l != nil {
				if tail == len(rays) {
					rays = append(rays, n.l)
				} else {
					rays[tail] = n.l
				}
				tail++
			}
			// branch down-right
			if n.r != nil {
				if tail == len(rays) {
					rays = append(rays, n.r)
				} else {
					rays[tail] = n.r
				}
				tail++
			}
			break
		}
	}

	return len(visited), nil

}
