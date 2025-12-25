package day04

import (
	"io"
)

func clearNeighborhood(grid field, width, height, x, y int) (removed int) {
	pos := y*width + x
	if grid[pos] == 0 {
		return 0
	}
	if !isRemoveable(grid, width, height, x, y) {
		return 0
	}

	// remove current
	grid[pos] = 0
	removed = 1

	// recurse to neighbors that might have become removable
	for neighbor := range neighbors {
		nx := x + neighbors[neighbor][0]
		ny := y + neighbors[neighbor][1]
		if nx < 0 || nx >= width || ny < 0 || ny >= height {
			continue
		}
		if grid[ny*width+nx] == 0 {
			continue
		}
		removed += clearNeighborhood(grid, width, height, nx, ny)
	}
	return removed
}

func Part2(in io.Reader) (int, error) {
	data, err := parseInput(in)
	if err != nil {
		return 0, err
	}

	totalRemoved := 0
	for pos := range data.field {
		if data.field[pos] == 0 {
			continue
		}
		x := pos % data.width
		y := pos / data.width
		totalRemoved += clearNeighborhood(data.field, data.width, data.row, x, y)
	}
	return totalRemoved, nil
}
