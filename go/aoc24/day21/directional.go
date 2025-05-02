package day21

import "github.com/donmahallem/aoc/aoc_utils"

var (
	DIRECTIONAL_LEFT  = Point{X: 0, Y: 1}
	DIRECTIONAL_RIGHT = Point{X: 2, Y: 1}
	DIRECTIONAL_UP    = Point{X: 1, Y: 0}
	DIRECTIONAL_DOWN  = Point{X: 1, Y: 1}
	DIRECTIONAL_A     = Point{X: 2, Y: 0}
)

func walkXFirstDirectional(dir *Point, currentDepth uint8, maxDepth uint8, cache *Cache) uint {
	var steps uint
	steps = 0
	current := &DIRECTIONAL_A // Stores current cursor
	var next *Point
	if dir.X != 0 {
		if dir.X < 0 {
			next = &DIRECTIONAL_LEFT
		} else {
			next = &DIRECTIONAL_RIGHT
		}
		for range aoc_utils.Abs(dir.X) {
			steps += WalkDirectional(current, next, currentDepth+1, maxDepth, cache)
			current = next
		}
	}
	if dir.Y != 0 {
		if dir.Y < 0 {
			next = &DIRECTIONAL_UP
		} else {
			next = &DIRECTIONAL_DOWN
		}
		for range aoc_utils.Abs(dir.Y) {
			steps += WalkDirectional(current, next, currentDepth+1, maxDepth, cache)
			current = next
		}
	}
	return steps + WalkDirectional(current, &DIRECTIONAL_A, currentDepth+1, maxDepth, cache)
}
func walkYFirstDirectional(dir *Point, currentDepth uint8, maxDepth uint8, cache *Cache) uint {
	var steps uint
	steps = 0
	current := &DIRECTIONAL_A // Stores current cursor
	var next *Point
	if dir.Y != 0 {
		if dir.Y < 0 {
			next = &DIRECTIONAL_UP
		} else {
			next = &DIRECTIONAL_DOWN
		}
		for range aoc_utils.Abs(dir.Y) {
			steps += WalkDirectional(current, next, currentDepth+1, maxDepth, cache)
			current = next
		}
	}
	if dir.X != 0 {
		if dir.X < 0 {
			next = &DIRECTIONAL_LEFT
		} else {
			next = &DIRECTIONAL_RIGHT
		}
		for range aoc_utils.Abs(dir.X) {
			steps += WalkDirectional(current, next, currentDepth+1, maxDepth, cache)
			current = next
		}
	}
	return steps + WalkDirectional(current, &DIRECTIONAL_A, currentDepth+1, maxDepth, cache)
}

func WalkDirectional(start *Point, end *Point, currentDepth uint8, maxDepth uint8, cache *Cache) uint {
	if currentDepth == maxDepth {
		return 1
	}
	cacheKey := HashId(start, end, maxDepth-currentDepth)
	if cachedValue, ok := (*cache)[cacheKey]; ok {
		return cachedValue
	}
	dir := start.Diff(*end)
	var returnValue uint
	if dir.X == 0 && dir.Y == 0 {
		// No walk necessary
		returnValue = WalkDirectional(&DIRECTIONAL_A, &DIRECTIONAL_A, currentDepth+1, maxDepth, cache)
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
