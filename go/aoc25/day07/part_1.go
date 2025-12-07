package day07

import (
	"io"
)

func Part1(in io.Reader) int {
	splitterMap, startX, startY, width, height := parseInput(in)

	rays := make([]int, 128)
	head, tail := 0, 1
	rays[0] = startY*width + startX

	visited := make(map[int]bool, width*height/4)
	totalCount := 0

	for head < tail {
		pos := rays[head]
		head++

		x := pos % width
		y := pos / width

		for {
			if visited[pos] {
				break
			}
			visited[pos] = true

			if y == height-1 {
				break
			}

			if _, ok := splitterMap[pos]; ok {
				totalCount++
				// branch down-left
				if x > 0 {
					nextPos := (y+1)*width + (x - 1)
					if !visited[nextPos] {
						if tail == len(rays) {
							rays = append(rays, nextPos)
						} else {
							rays[tail] = nextPos
						}
						tail++
					}
				}
				// branch down-right
				if x < width-1 {
					nextPos := (y+1)*width + (x + 1)
					if !visited[nextPos] {
						if tail == len(rays) {
							rays = append(rays, nextPos)
						} else {
							rays[tail] = nextPos
						}
						tail++
					}
				}
				break
			}

			// move straight down
			pos += width
			y++
			// x stays the same
		}
	}

	return totalCount
}
