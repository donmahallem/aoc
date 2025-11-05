package day06

import "io"

func checkForLoop(field *Field, guard Guard, visited []int32, token int32) bool {
	for {
		stateIdx := int32(translateToIndex(guard.x, guard.y, field.width))*4 + int32(guard.dir)
		if visited[stateIdx] == token {
			return true
		}
		visited[stateIdx] = token

		switch guard.dir {
		case DIR_DOWN:
			if guard.y == 0 {
				return false
			}
			nextY := guard.y - 1
			if field.HasObstacle(guard.x, nextY) {
				guard.dir = DIR_RIGHT
				continue
			}
			guard.y = nextY

		case DIR_UP:
			if guard.y+1 >= field.height {
				return false
			}
			nextY := guard.y + 1
			if field.HasObstacle(guard.x, nextY) {
				guard.dir = DIR_LEFT
				continue
			}
			guard.y = nextY

		case DIR_RIGHT:
			if guard.x+1 >= field.width {
				return false
			}
			nextX := guard.x + 1
			if field.HasObstacle(nextX, guard.y) {
				guard.dir = DIR_UP
				continue
			}
			guard.x = nextX

		case DIR_LEFT:
			if guard.x == 0 {
				return false
			}
			nextX := guard.x - 1
			if field.HasObstacle(nextX, guard.y) {
				guard.dir = DIR_DOWN
				continue
			}
			guard.x = nextX
		}
	}
}

func translateFromIndex(idx, width uint16) (uint16, uint16) {
	return idx % width, idx / width
}

func Part2(in io.Reader) int {
	field, guard, err := ReadSource(in)
	if err != nil {
		panic(err)
	}
	pathInfo := leaveArea(&field, guard)
	startIdx := translateToIndex(guard.x, guard.y, field.width)

	stateSlots := make([]int32, int(field.width)*int(field.height)*4)
	token := int32(1)

	blockages := 0
	for _, step := range pathInfo.Steps {
		if step.Index == startIdx {
			continue
		}

		if token == 0 {
			for i := range stateSlots {
				stateSlots[i] = 0
			}
			token = 1
		}

		field.obstacles[step.Index] = struct{}{}
		if checkForLoop(&field, step.PreGuard, stateSlots, token) {
			blockages++
		}
		delete(field.obstacles, step.Index)
		token++
	}

	return blockages
}
