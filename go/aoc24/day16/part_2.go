package day16

import (
	"io"
)

type visitedMap = map[point]bool
type traceBackState struct {
	// Position in the grid where you are
	pos point
	// Value of cell where you are
	value int
	// Value of cell you are coming from
	previousValue int
	// Direction you walked from the last cell
	dir *direction
}

// Takes a state and checks neighbours for cells that could have been walked form start to finish
func walkBackPathValues(field *field, traceState *traceBackState, visited *visitedMap) bool {
	if traceState.value == 0 {
		// Endpoint reached
		(*visited)[traceState.pos] = true
		return true
	}
	var diff, nextValue int
	status := false
	var testCoord point = point{}
	testDirs := [3]*direction{traceState.dir, translateLeft(traceState.dir), translateRight(traceState.dir)}
	for testDirIdx, testDir := range testDirs {
		if testDir == nil {
			continue
		}
		testCoord.X = traceState.pos.X + testDir.X
		testCoord.Y = traceState.pos.Y + testDir.Y
		// bounds check
		tx := int(testCoord.X)
		ty := int(testCoord.Y)
		if ty < 0 || ty >= len(*field) || tx < 0 || tx >= len((*field)[ty]) {
			continue
		}
		nextValue = (*field)[ty][tx]
		if nextValue == CELL_WALL {
			continue
		}
		diff = traceState.value - nextValue
		if diff == 1 || diff == 1001 {
			state := traceBackState{pos: testCoord,
				value:         nextValue,
				dir:           testDir,
				previousValue: traceState.value}
			status = walkBackPathValues(field, &state, visited) || status
		} else if testDirIdx == 0 {
			// Check overlapping routes
			diff = traceState.previousValue - nextValue
			if diff == 2 {
				state := traceBackState{pos: testCoord,
					value:         nextValue,
					dir:           testDir,
					previousValue: traceState.previousValue - 1}
				status = walkBackPathValues(field, &state, visited) || status
				continue
			}
			// Check interleaving routes
			ntx := tx + int(testDir.X)
			nty := ty + int(testDir.Y)
			if nty < 0 || nty >= len(*field) || ntx < 0 || ntx >= len((*field)[nty]) {
				continue
			}
			nextValue = (*field)[nty][ntx]
			if nextValue != CELL_WALL && traceState.value-nextValue == 2 {
				state := traceBackState{pos: testCoord,
					value:         traceState.value - 1,
					dir:           testDir,
					previousValue: traceState.value}
				status = walkBackPathValues(field, &state, visited) || status
			}
		}
	}
	if status {
		(*visited)[traceState.pos] = true
	}
	return status
}

func findVisitedCells(field *field, start *point, end *point) int {
	visited := make(visitedMap)
	var testCoord point = point{}
	var currentCellValue, nextCellValue, diff int
	// bounds check
	if field == nil || len(*field) == 0 {
		return 0
	}
	if int(end.Y) < 0 || int(end.Y) >= len(*field) || int(end.X) < 0 || int(end.X) >= len((*field)[end.Y]) {
		return 0
	}
	currentCellValue = (*field)[end.Y][end.X]
	status := false
	for _, dir := range dirsALL {
		testCoord.X = end.X + dir.X
		testCoord.Y = end.Y + dir.Y
		tx := int(testCoord.X)
		ty := int(testCoord.Y)
		if ty < 0 || ty >= len(*field) || tx < 0 || tx >= len((*field)[ty]) {
			continue
		}
		nextCellValue = (*field)[ty][tx]
		if nextCellValue == CELL_WALL {
			continue
		}
		diff = currentCellValue - nextCellValue
		if diff == 1001 || diff == 1 {
			tr := traceBackState{value: nextCellValue, previousValue: currentCellValue, pos: testCoord, dir: &dir}
			status = walkBackPathValues(field, &tr, &visited) || status
		}
	}
	if status {
		visited[*end] = true
	}
	return len(visited)
}

func Part2(in io.Reader) (int, error) {
	data, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	calculatePathValues(&data.Field, &data.Start)
	return findVisitedCells(&data.Field, &data.Start, &data.End), nil
}
