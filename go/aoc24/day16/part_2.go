package day16

import (
	"fmt"
	"io"
	"slices"
)

func TakePath(field *Field, takenPath *[]Point, currentDir Direction, currentValue uint, maxValue uint, takenCells *map[Point]bool) bool {
	initalPathLength := len(*takenPath)
	currentCoord := (*takenPath)[initalPathLength-1]
	if currentCoord.X == 13 && currentCoord.Y == 1 {
		fmt.Printf("Val: %d\n", currentValue)
		(*takenCells)[currentCoord] = true
		return true
	}
	checkDirs := make([]Direction, 0, 3)
	checkDirs = append(checkDirs, currentDir, *translateLeft(&currentDir), *translateRight(&currentDir))
	nextCoord := Point{}
	valid := false
	for checkDirIdx, checkDir := range checkDirs {
		nextCoord.X = currentCoord.X + checkDir.X
		nextCoord.Y = currentCoord.Y + checkDir.Y
		if (*field)[nextCoord.Y][nextCoord.X] == CELL_WALL || slices.Contains(*takenPath, nextCoord) {
			//fmt.Printf("Wall at %v\n", nextCoord)
			continue
		}
		nextValue := currentValue + 1
		// Idx 0 is straight ahead, all others are turned
		if checkDirIdx > 0 {
			nextValue += 1000
		}
		*takenPath = append(*takenPath, nextCoord)
		valid = valid || TakePath(field, takenPath, checkDir, nextValue, maxValue, takenCells)
		if valid {
			(*takenCells)[currentCoord] = true
		}
		*takenPath = (*takenPath)[0 : initalPathLength-1]
	}
	return valid
}

func CountShortestPathCells2(field *Field, start *Point, end *Point, pathValue uint) uint {
	initialPath := make([]Point, 0)
	takenCells := make(map[Point]bool)
	initialPath = append(initialPath, *start)
	TakePath(field, &initialPath, DIR_RIGHT, 0, pathValue, &takenCells)
	fmt.Printf("%v", takenCells)
	return uint(len(takenCells))
}

func CountShortestPathCells(field *Field, start *Point, end *Point, pathValue uint) uint {
	var walk func(Point, uint, Direction) bool
	takenCells := make(map[Point]bool)
	checkDirs := make([]Direction, 0, 3)
	walk = func(currentCoord Point, currentValue uint, dir Direction) bool {
		//printField(field, &currentCoord, dir)
		if currentCoord == *end {
			takenCells[currentCoord] = true
			return true
		} else if currentValue > pathValue {
			return false
		}
		var nextCoord Point = Point{}
		status := false
		checkDirs = append(checkDirs[:0], dir, *translateLeft(&dir), *translateRight(&dir))
		for checkDirIdx, checkDir := range checkDirs {
			nextCoord.X = currentCoord.X + checkDir.X
			nextCoord.Y = currentCoord.Y + checkDir.Y
			if (*field)[nextCoord.Y][nextCoord.X] == CELL_WALL {
				//fmt.Printf("Wall at %v\n", nextCoord)
				continue
			}
			nextValue := currentValue + 1
			// Idx 0 is straight ahead, all others are turned
			if checkDirIdx > 0 {
				nextValue += 1000
			}
			status = status || walk(nextCoord, nextValue, checkDir)
		}
		if status {
			takenCells[currentCoord] = true
		}
		return status
	}
	walk(*start, 0, DIR_RIGHT)
	fmt.Printf("%v\n", takenCells)
	return uint(len(takenCells))
}

func Part2(in io.Reader) int {
	field, start, end := ParseInput(in)
	pathValue := FindShortestPath(&field, &start, &end)

	return int(CountShortestPathCells2(&field, &start, &end, pathValue))
}
