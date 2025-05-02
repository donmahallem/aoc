package day21

var numericLookup map[byte]Point = map[byte]Point{
	'7': {X: 0, Y: 0},
	'8': {X: 1, Y: 0},
	'9': {X: 2, Y: 0},
	'4': {X: 0, Y: 1},
	'5': {X: 1, Y: 1},
	'6': {X: 2, Y: 1},
	'1': {X: 0, Y: 2},
	'2': {X: 1, Y: 2},
	'3': {X: 2, Y: 2},
	'0': {X: 1, Y: 3},
	'A': {X: 2, Y: 3},
}

var (
	NUMERIC_0 = Point{X: 1, Y: 3}
	NUMERIC_1 = Point{X: 0, Y: 2}
	NUMERIC_2 = Point{X: 1, Y: 2}
	NUMERIC_3 = Point{X: 2, Y: 2}
	NUMERIC_4 = Point{X: 0, Y: 1}
	NUMERIC_5 = Point{X: 1, Y: 1}
	NUMERIC_6 = Point{X: 2, Y: 1}
	NUMERIC_7 = Point{X: 0, Y: 0}
	NUMERIC_8 = Point{X: 1, Y: 0}
	NUMERIC_9 = Point{X: 2, Y: 0}
	NUMERIC_A = Point{X: 2, Y: 3}
)

// Helper function for WalkNumericSequence
func walkNumericSequenceSub(start *Point, end *Point, currentDepth uint8, maxDepth uint8, cache *Cache) uint {
	dir := start.Diff(*end)
	if dir.X == 0 && dir.Y == 0 {
		// No walking necessary
		return WalkDirectional(&DIRECTIONAL_A, &DIRECTIONAL_A, currentDepth+1, maxDepth, cache)
	}
	if start.X == 0 && end.Y == 3 {
		// First walk X than Y
		// Relevant for moves from 1,4,7 to 0,A
		return walkXFirstDirectional(dir, currentDepth, maxDepth, cache)
	} else if end.X == 0 && start.Y == 3 {
		// First walk Y than X
		return walkYFirstDirectional(dir, currentDepth, maxDepth, cache)
	} else {
		// try first x and first y
		endResultA := walkXFirstDirectional(dir, currentDepth, maxDepth, cache)
		if endResultB := walkYFirstDirectional(dir, currentDepth, maxDepth, cache); endResultB < endResultA {
			return endResultB
		}
		return endResultA
	}
}

// Iterates over the given Byte-Sequence from the Numeric pad to a depth of maxDepth (inclusive)
func WalkNumericSequence(sequence *[]byte, maxDepth uint8, cache *Cache) uint {
	var pathCost uint = 0
	var previousLetter byte = 'A'
	for _, currentLetter := range *sequence {
		from := numericLookup[previousLetter]
		to := numericLookup[currentLetter]
		pathCost += walkNumericSequenceSub(&from, &to, 0, maxDepth, cache)
		previousLetter = currentLetter
	}
	return pathCost
}
