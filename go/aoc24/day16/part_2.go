package day16

import (
	"fmt"
	"io"
)

type VisitedMap = map[Point]bool
type TraceBackState struct {
	// Position in the grid where you are
	pos Point
	// Value of cell where you are
	value int
	// Value of cell you are coming from
	previousValue int
	// Direction you walked from the last cell
	dir *Direction
}

// Takes a state and checks neighbours for cells that could have been walked form start to finish
func walkBackPathValues(field *Field, traceState *TraceBackState, visited *VisitedMap) bool {
	fmt.Printf("Tested: %v\n", traceState.pos)
	if traceState.value == 0 {
		// Endpoint reached
		(*visited)[traceState.pos] = true
		return true
	}
	var diff, nextValue int
	status := false
	var testCoord Point = Point{}
	testDirs := [3]*Direction{traceState.dir, translateLeft(traceState.dir), translateRight(traceState.dir)}
	for testDirIdx, testDir := range testDirs {
		testCoord.X = traceState.pos.X + testDir.X
		testCoord.Y = traceState.pos.Y + testDir.Y
		nextValue = (*field)[testCoord.Y][testCoord.X]
		if nextValue == CELL_WALL {
			continue
		}
		diff = traceState.value - nextValue
		if diff == 1 || diff == 1001 {
			state := TraceBackState{pos: testCoord,
				value:         nextValue,
				dir:           testDir,
				previousValue: traceState.value}
			status = walkBackPathValues(field, &state, visited) || status
		} else if testDirIdx == 0 {
			// Check overlapping routes
			diff = traceState.previousValue - nextValue
			if diff == 2 {
				state := TraceBackState{pos: testCoord,
					value:         nextValue,
					dir:           testDir,
					previousValue: traceState.previousValue - 1}
				status = walkBackPathValues(field, &state, visited) || status
				continue
			}
			// Check interleaving routes
			nextValue = (*field)[testCoord.Y+testDir.Y][testCoord.X+testDir.X]
			if nextValue != CELL_WALL && traceState.value-nextValue == 2 {
				state := TraceBackState{pos: testCoord,
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

func FindVisitedCells(field *Field, start *Point, end *Point) int {
	visited := make(VisitedMap)
	var testCoord Point = Point{}
	var currentCellValue, nextCellValue, diff int
	currentCellValue = (*field)[end.Y][end.X]
	status := false
	for _, dir := range DIRS_ALL {
		testCoord.X = end.X + dir.X
		testCoord.Y = end.Y + dir.Y
		nextCellValue = (*field)[testCoord.Y][testCoord.X]
		if nextCellValue == CELL_WALL {
			continue
		}
		diff = currentCellValue - nextCellValue
		if diff == 1001 || diff == 1 {
			tr := TraceBackState{value: nextCellValue, previousValue: currentCellValue, pos: testCoord, dir: &dir}
			status = walkBackPathValues(field, &tr, &visited) || status
		}
	}
	if status {
		visited[*end] = true
	}
	return len(visited)
}

func Part2(in io.Reader) int {
	field, start, end := ParseInput(in)
	CalculatePathValues(&field, &start)
	return FindVisitedCells(&field, &start, &end)
}
