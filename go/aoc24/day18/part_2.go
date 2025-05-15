package day18

import (
	"fmt"
	"io"
)

func IsPathAvailable(field Field, pointIdx uint16, fieldWidth uint16, fieldHeight uint16) bool {
	fieldWidthInt := int16(fieldWidth)
	fieldHeightInt := int16(fieldHeight)
	queue := []Point{{X: 0, Y: 0}}
	visited := make(map[Point]bool, 64)
	var currentPosition Point
	for len(queue) > 0 {
		currentPosition = queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		visited[currentPosition] = true

		for _, checkDir := range DIRS_ALL {
			nextPoint := Point{X: currentPosition.X + checkDir.X, Y: currentPosition.Y + checkDir.Y}
			if visited[nextPoint] {
				continue
			}
			if nextPoint.X < 0 || nextPoint.Y < 0 || nextPoint.X >= fieldWidthInt || nextPoint.Y >= fieldHeightInt {
				// next coord outside the field dimensions
				continue
			}
			currentCellValue := field[nextPoint.Y][nextPoint.X]
			if currentCellValue > 0 && currentCellValue <= pointIdx {
				// Cell Corrupted
				continue
			}
			if nextPoint.X == fieldWidthInt-1 && nextPoint.Y == fieldHeightInt-1 {
				return true
			}
			queue = append(queue, nextPoint)
		}
	}
	return false
}

func FindFirstNonSolvable(points Field, maxStep uint16, fieldWidth uint16, fieldHeight uint16) uint16 {
	var left uint16 = 0
	right := maxStep
	for left < right-1 {
		mid := (left + right) / 2
		ok := IsPathAvailable(points, mid, fieldWidth, fieldHeight)
		if ok {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

func Part2(in io.Reader) Point {
	points := ParseInput(in)
	field := PointsToField(points, 71, 71)
	result := FindFirstNonSolvable(field, uint16(len(points)), 71, 71)
	fmt.Printf("%v\n", [2]int16{points[result].X, points[result].Y})
	return points[result]
}
