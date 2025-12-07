package day07

import (
	"io"
)

func Part1(in io.Reader) int {
	splitterMap, startX, startY, width, height := parseInput(in)

	rays := make([]int, 128)
	head, tail := 0, 1
	rays[0] = startY*width + startX

	visited := make(map[int]bool, width*height/2)
	totalCount := 0

	for head < tail {
		pos := rays[tail-1]
		tail--

		if visited[pos] {
			continue
		}
		visited[pos] = true

		x := pos % width
		y := pos / width
		if y == height-1 {
			continue
		}

		if _, ok := splitterMap[pos]; !ok {
			nextPos := pos + width
			if tail == len(rays) {
				rays = append(rays, nextPos)
			} else {
				rays[tail] = nextPos
			}
			tail++
			continue
		}

		totalCount++
		if x > 0 {
			nextPos := (y+1)*width + (x - 1)
			if tail == len(rays) {
				rays = append(rays, nextPos)
			} else {
				rays[tail] = nextPos
			}
			tail++
		}
		if x < width-1 {
			nextPos := (y+1)*width + (x + 1)
			if tail == len(rays) {
				rays = append(rays, nextPos)
			} else {
				rays[tail] = nextPos
			}
			tail++
		}
	}

	return totalCount
}
