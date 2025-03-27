package day14

import (
	"fmt"
	"os"
)

func SortHorizontal(a, b Robot) int {
	if a.pos.y == b.pos.y {
		return a.pos.x - b.pos.x
	}
	return a.pos.y - b.pos.y
}

func Step(robots *[]Robot, width *int, height *int) {
	for i := range *robots {
		(*robots)[i].pos.x = ((*robots)[i].pos.x + (*robots)[i].vec.x) % *width
		(*robots)[i].pos.y = ((*robots)[i].pos.y + (*robots)[i].vec.y) % *height
		if (*robots)[i].pos.x < 0 {
			(*robots)[i].pos.x = *width + (*robots)[i].pos.x
		}
		if (*robots)[i].pos.y < 0 {
			(*robots)[i].pos.y = *height + (*robots)[i].pos.y
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

func Part2(in *os.File) {
	data := LoadFile(in)
	width, height, maxDepth := 101, 103, 10000000
	totalSum := FindNonDouble(&data, maxDepth, &width, &height)
	fmt.Printf("%d\n", maxDepth-totalSum)
}
