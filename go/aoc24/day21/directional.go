package day21

import "github.com/donmahallem/aoc/go/aoc_utils/int_util"

var (
	directionalPosition_Left  = Point{X: 0, Y: 1}
	directionalPosition_Right = Point{X: 2, Y: 1}
	directionalPosition_Up    = Point{X: 1, Y: 0}
	directionalPosition_Down  = Point{X: 1, Y: 1}
	directionalPosition_A     = Point{X: 2, Y: 0}
)

func walkXFirstDirectional(dir *Point, currentDepth uint8, maxDepth uint8, cache *cache) uint {
	var steps uint
	steps = 0
	current := &directionalPosition_A // Stores current cursor
	var next *Point
	if dir.X != 0 {
		if dir.X < 0 {
			next = &directionalPosition_Left
		} else {
			next = &directionalPosition_Right
		}
		for range int_util.AbsInt(dir.X) {
			steps += WalkDirectional(current, next, currentDepth+1, maxDepth, cache)
			current = next
		}
	}
	if dir.Y != 0 {
		if dir.Y < 0 {
			next = &directionalPosition_Up
		} else {
			next = &directionalPosition_Down
		}
		for range int_util.AbsInt(dir.Y) {
			steps += WalkDirectional(current, next, currentDepth+1, maxDepth, cache)
			current = next
		}
	}
	return steps + WalkDirectional(current, &directionalPosition_A, currentDepth+1, maxDepth, cache)
}
func walkYFirstDirectional(dir *Point, currentDepth uint8, maxDepth uint8, cache *cache) uint {
	var steps uint
	steps = 0
	current := &directionalPosition_A // Stores current cursor
	var next *Point
	if dir.Y != 0 {
		if dir.Y < 0 {
			next = &directionalPosition_Up
		} else {
			next = &directionalPosition_Down
		}
		for range int_util.AbsInt(dir.Y) {
			steps += WalkDirectional(current, next, currentDepth+1, maxDepth, cache)
			current = next
		}
	}
	if dir.X != 0 {
		if dir.X < 0 {
			next = &directionalPosition_Left
		} else {
			next = &directionalPosition_Right
		}
		for range int_util.AbsInt(dir.X) {
			steps += WalkDirectional(current, next, currentDepth+1, maxDepth, cache)
			current = next
		}
	}
	return steps + WalkDirectional(current, &directionalPosition_A, currentDepth+1, maxDepth, cache)
}

func WalkDirectional(start *Point, end *Point, currentDepth uint8, maxDepth uint8, cache *cache) uint {
	if currentDepth == maxDepth {
		return 1
	}
	cacheKey := hashId(start, end, maxDepth-currentDepth)
	if cachedValue, ok := (*cache)[cacheKey]; ok {
		return cachedValue
	}
	dir := start.Diff(*end)
	var returnValue uint
	if dir.X == 0 && dir.Y == 0 {
		// No walk necessary
		returnValue = WalkDirectional(&directionalPosition_A, &directionalPosition_A, currentDepth+1, maxDepth, cache)
	} else if start.X == 0 && end.Y == 0 {
		// First walk X than Y
		returnValue = walkXFirstDirectional(dir, currentDepth, maxDepth, cache)
	} else if start.Y == 0 && end.X == 0 {
		// First walk Y than X
		returnValue = walkYFirstDirectional(dir, currentDepth, maxDepth, cache)
	} else {
		// try first x and first y
		resultWalkXFirst := walkXFirstDirectional(dir, currentDepth, maxDepth, cache)
		if resultWalkYFirst := walkYFirstDirectional(dir, currentDepth, maxDepth, cache); resultWalkYFirst < resultWalkXFirst {
			returnValue = resultWalkYFirst
		} else {
			returnValue = resultWalkXFirst
		}
	}
	(*cache)[cacheKey] = returnValue
	return returnValue
}
