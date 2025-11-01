package day14

import (
	"io"
)

func SortHorizontal(a, b Robot) int {
	if a.pos.Y == b.pos.Y {
		return a.pos.X - b.pos.X
	}
	return a.pos.Y - b.pos.Y
}

func Step(robots []Robot, width int, height int) {
	for i := range robots {
		robot := &robots[i]
		robot.pos.X = (robot.pos.X + robot.vec.X) % width
		robot.pos.Y = (robot.pos.Y + robot.vec.Y) % height
		if robot.pos.X < 0 {
			robot.pos.X = width + robot.pos.X
		}
		if robot.pos.Y < 0 {
			robot.pos.Y = height + robot.pos.Y
		}
	}
}

func FindNonDouble(robots []Robot, maxDepth int, width int, height int) int {
	if maxDepth < 0 {
		return -1
	}
	fieldSize := width * height
	if fieldSize <= 0 {
		return -1
	}
	visited := make([]int, fieldSize)
	stepID := 0
	for remaining := maxDepth; remaining >= 0; remaining-- {
		Step(robots, width, height)
		stepID++
		duplicateFound := false
		for i := range robots {
			pos := robots[i].pos
			idx := pos.Y*width + pos.X
			if visited[idx] == stepID {
				duplicateFound = true
				break
			}
			visited[idx] = stepID
		}
		if !duplicateFound {
			return remaining - 1
		}
	}
	return -1
}

func Part2(in io.Reader) int {
	data := LoadFile(in)
	width, height, maxDepth := 101, 103, 10000000
	totalSum := FindNonDouble(data, maxDepth, width, height)
	return maxDepth - totalSum
}
