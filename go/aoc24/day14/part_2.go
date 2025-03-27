package day14

import (
	"fmt"
	"io"
)

func SortHorizontal(a, b Robot) int {
	if a.pos.Y == b.pos.Y {
		return a.pos.X - b.pos.X
	}
	return a.pos.Y - b.pos.Y
}

func Step(robots *[]Robot, width *int, height *int) {
	for i := range *robots {
		(*robots)[i].pos.X = ((*robots)[i].pos.X + (*robots)[i].vec.X) % *width
		(*robots)[i].pos.Y = ((*robots)[i].pos.Y + (*robots)[i].vec.Y) % *height
		if (*robots)[i].pos.X < 0 {
			(*robots)[i].pos.X = *width + (*robots)[i].pos.X
		}
		if (*robots)[i].pos.Y < 0 {
			(*robots)[i].pos.Y = *height + (*robots)[i].pos.Y
		}
	}
}

func FindNonDouble(robots *[]Robot, maxDepth int, width *int, height *int) int {
	if maxDepth < 0 {
		return -1
	}
	Step(robots, width, height)
	numRobots := len(*robots)
	duplicateFound := false
	for i := 0; i < numRobots-1 && !duplicateFound; i++ {
		for j := i + 1; j < numRobots; j++ {
			if (*robots)[i].pos == (*robots)[j].pos {
				duplicateFound = true
				break
			}
		}
	}
	if !duplicateFound {
		return maxDepth - 1
	}
	return FindNonDouble(robots, maxDepth-1, width, height)
}

func Part2(in io.Reader) {
	data := LoadFile(in)
	width, height, maxDepth := 101, 103, 10000000
	totalSum := FindNonDouble(&data, maxDepth, &width, &height)
	fmt.Printf("%d\n", maxDepth-totalSum)
}
