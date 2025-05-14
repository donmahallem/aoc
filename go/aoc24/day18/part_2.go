package day18

import (
	"io"
)

func IsPathAvailable(field Field, scoreField Field, pointIdx int, fieldWidth int, fieldHeight int) bool {
	checkPositions := []Point{{X: 0, Y: 0}}
	var currentPosition Point
	nextCoord := Point{}
	for len(checkPositions) > 0 {
		currentPosition = checkPositions[0]
		checkPositions = checkPositions[1:]

		for _, checkDir := range DIRS_ALL {
			nextCoord.X = currentPosition.X + checkDir.X
			nextCoord.Y = currentPosition.Y + checkDir.Y
			if nextCoord.X < 0 || nextCoord.Y < 0 || nextCoord.X >= fieldWidth || nextCoord.Y >= fieldHeight {
				// next coord outside the field dimensions
				continue
			} else if field[nextCoord.Y][nextCoord.X] == CELL_CORRUPTED {
				// Cell is corrupted
				continue
			} else if scoreField[nextCoord.Y][nextCoord.X] == pointIdx {
				// Cell already visited
				continue
			}
			scoreField[nextCoord.Y][nextCoord.X] = pointIdx
			if nextCoord.X == fieldWidth-1 && nextCoord.Y == fieldHeight-1 {
				return true
			}
			checkPositions = append(checkPositions, nextCoord)
		}
	}
	return false
}

func FindFirstNonSolvable(points []Point, fieldWidth, fieldHeight int) (*Point, bool) {
	obstacleField := CreateEmptyField(uint(fieldWidth), uint(fieldHeight))
	scoreField := CreateEmptyField(uint(fieldWidth), uint(fieldHeight))
	lastIdx := -1
	var ok bool
	for pointIdx, point := range points {
		obstacleField[point.Y][point.X] = CELL_CORRUPTED
		if lastIdx >= 0 && scoreField[point.Y][point.X] < lastIdx {
			continue
		}
		ok = IsPathAvailable(obstacleField, scoreField, pointIdx+1, fieldWidth, fieldHeight)
		if ok {
			lastIdx = pointIdx
		} else {
			return &point, true
		}
	}
	return nil, false
}

func Part2(in io.Reader) [2]int {
	points := ParseInput(in)
	result, ok := FindFirstNonSolvable(points, 71, 71)
	if ok {
		return [2]int{result.X, result.Y}
	} else {
		return [2]int{-1, -1}
	}
}
