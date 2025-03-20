package day06

import (
	"fmt"
	"os"
)

func testLoop(field *Field, guard Guard) bool {
	walked := make(map[[3]int]bool)
	for {
		if guard.dir == DIR_DOWN {
			nextY := guard.y - 1
			key := [3]int{nextY, guard.x, guard.dir}
			if OutOfBounds(field, &guard.x, &nextY) {
				return false
			} else if (*field).field[nextY][guard.x] {
				guard.dir = DIR_RIGHT
			} else if walked[key] {
				return true
			} else {
				guard.y = nextY
				walked[key] = true
			}
		} else if guard.dir == DIR_UP {
			nextY := guard.y + 1
			key := [3]int{nextY, guard.x, guard.dir}
			if OutOfBounds(field, &guard.x, &nextY) {
				return false
			} else if (*field).field[nextY][guard.x] {
				guard.dir = DIR_LEFT
			} else if walked[key] {
				return true
			} else {
				guard.y = nextY
				walked[key] = true
			}
		} else if guard.dir == DIR_RIGHT {
			nextX := guard.x + 1
			key := [3]int{guard.y, nextX, guard.dir}
			if OutOfBounds(field, &nextX, &guard.y) {
				return false
			} else if (*field).field[guard.y][nextX] {
				guard.dir = DIR_UP
			} else if walked[key] {
				return true
			} else {
				guard.x = nextX
				walked[key] = true
			}
		} else if guard.dir == DIR_LEFT {
			nextX := guard.x - 1
			key := [3]int{guard.y, nextX, guard.dir}
			if OutOfBounds(field, &nextX, &guard.y) {
				return false
			} else if (*field).field[guard.y][nextX] {
				guard.dir = DIR_DOWN
			} else if walked[key] {
				return true
			} else {
				guard.x = nextX
				walked[key] = true
			}
		}
	}
}
func Part2() {
	obstacles, guard, err := readSource(os.Stdin)
	if err != nil {
		panic(err)
	}
	basePath := leaveArea(&obstacles, guard)
	blockages := 0
	for key := range basePath {
		if key[0] == guard.y && key[1] == guard.x {
			// skip as this is the start position and can't be blocked
			continue
		}
		// set temporary obstacle
		obstacles.field[key[0]][key[1]] = true
		// Check for loop
		if testLoop(&obstacles, guard) {
			blockages++
		}
		// remove temporary obstacle
		obstacles.field[key[0]][key[1]] = false
	}
	fmt.Printf("%d\n", blockages)
}
