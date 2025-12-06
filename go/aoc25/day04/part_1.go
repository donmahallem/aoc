package day04

import (
	"io"
)

var neighbors [8][2]int = [8][2]int{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func isRemoveable(f field, width, height, x, y int) bool {
	pos := y*width + x
	if f[pos] == 0 {
		return false
	}

	count := 0
	for _, d := range neighbors {
		nx, ny := x+d[0], y+d[1]
		if nx < 0 || nx >= width || ny < 0 || ny >= height {
			continue
		}
		if f[ny*width+nx] == 1 {
			count++
		}
	}
	return count < 4
}

func Part1(in io.Reader) int {
	grid, width, height := parseInput(in)

	totalCount := 0

	for y := range height {
		for x := range width {
			if isRemoveable(grid, width, height, x, y) {
				totalCount++
			}
		}
	}
	return totalCount
}
