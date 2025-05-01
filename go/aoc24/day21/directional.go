package day21

import "github.com/donmahallem/aoc/aoc_utils"

var (
	DIRECTIONAL_LEFT  = Point{X: 0, Y: 1}
	DIRECTIONAL_RIGHT = Point{X: 2, Y: 1}
	DIRECTIONAL_UP    = Point{X: 1, Y: 0}
	DIRECTIONAL_DOWN  = Point{X: 1, Y: 1}
	DIRECTIONAL_A     = Point{X: 2, Y: 0}
)

func walkXFirstDirectional(dir *Point, currentDepth uint, maxDepth uint) uint {
	var steps, tempSteps uint
	steps = 0
	current := &DIRECTIONAL_A /// ERSETZE MIT A WEIL START NICHT START IST
	var next *Point
	if dir.X != 0 {
		if dir.X < 0 {
			next = &DIRECTIONAL_LEFT
		} else {
			next = &DIRECTIONAL_RIGHT
		}
		for range aoc_utils.Abs(dir.X) {
			steps += WalkDirectional(current, next, currentDepth+1, maxDepth)
			current = next
		}
	}
	if dir.Y == 0 {
		// Fast exit
		// No Y moves
		tempSteps = WalkDirectional(current, &DIRECTIONAL_A, currentDepth+1, maxDepth)
		return steps + tempSteps
	} else if dir.Y < 0 {
		next = &DIRECTIONAL_UP
	} else {
		next = &DIRECTIONAL_DOWN
	}
	for range aoc_utils.Abs(dir.Y) {
		steps += WalkDirectional(current, next, currentDepth+1, maxDepth)
		current = next
	}
	tempSteps = WalkDirectional(next, &DIRECTIONAL_A, currentDepth+1, maxDepth)
	steps += tempSteps
	return steps
}
func walkYFirstDirectional(dir *Point, currentDepth uint, maxDepth uint) uint {
	var steps, tempSteps uint
	steps = 0
	current := &DIRECTIONAL_A
	var next *Point
	if dir.Y != 0 {
		if dir.Y < 0 {
			next = &DIRECTIONAL_UP
		} else {
			next = &DIRECTIONAL_DOWN
		}
		for range aoc_utils.Abs(dir.Y) {
			steps += WalkDirectional(current, next, currentDepth+1, maxDepth)
			current = next
		}
	}
	if dir.X == 0 {
		// Fast exit
		// No Y moves
		tempSteps = WalkDirectional(current, &DIRECTIONAL_A, currentDepth+1, maxDepth)
		return steps + tempSteps
	} else if dir.X < 0 {
		next = &DIRECTIONAL_LEFT
	} else {
		next = &DIRECTIONAL_RIGHT
	}
	for range aoc_utils.Abs(dir.X) {
		steps += WalkDirectional(current, next, currentDepth+1, maxDepth)
		current = next
	}
	steps += WalkDirectional(next, &DIRECTIONAL_A, currentDepth+1, maxDepth)
	return steps
}

func WalkDirectional(start *Point, end *Point, currentDepth uint, maxDepth uint) uint {
	if currentDepth == maxDepth {
		return 1
	}
	dir := start.Diff(*end)
	if dir.X == 0 && dir.Y == 0 {
		// No walk necessary
		return WalkDirectional(&DIRECTIONAL_A, &DIRECTIONAL_A, currentDepth+1, maxDepth)
	}
	if start.X == 0 && end.Y == 0 {
		// First walk X than Y
		return walkXFirstDirectional(dir, currentDepth, maxDepth)
	} else if end.Y == 0 && start.X == 0 {
		// First walk Y than X
		return walkYFirstDirectional(dir, currentDepth, maxDepth)
	} else {
		// try first x and first y
		endResultA := walkXFirstDirectional(dir, currentDepth, maxDepth)
		if endResultB := walkYFirstDirectional(dir, currentDepth, maxDepth); endResultB < endResultA {
			return endResultB
		}
		return endResultA
	}
}
