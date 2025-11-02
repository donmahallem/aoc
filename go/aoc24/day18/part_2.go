package day18

import (
	"io"

	"container/list"
)

func IsPathAvailable(field Field, pointIdx, fieldWidth, fieldHeight int16) bool {
	fieldWidthInt := fieldWidth
	fieldHeightInt := fieldHeight
	queue := list.New()
	queue.PushBack(Point{X: 0, Y: 0})
	visited := make(map[Point]bool, 64)
	var currentPosition Point
	for queue.Len() > 0 {
		currentPosition = queue.Remove(queue.Back()).(Point)

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
			idx := nextPoint.Y*fieldWidthInt + nextPoint.X
			currentCellValue := field[idx]
			if currentCellValue > 0 && currentCellValue <= pointIdx {
				// Cell Corrupted
				continue
			}
			if nextPoint.X == fieldWidthInt-1 && nextPoint.Y == fieldHeightInt-1 {
				return true
			}
			queue.PushBack(nextPoint)
		}
	}
	return false
}

func FindFirstNonSolvable(field Field, maxStep, fieldWidth, fieldHeight int16) int16 {
	var left int16 = 0
	right := maxStep
	for left < right-1 {
		mid := (left + right) / 2
		ok := IsPathAvailable(field, mid, fieldWidth, fieldHeight)
		if ok {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

func Part2Base(in io.Reader, width, height int16) Point {
	parsedData := ParseInput(in, width, height)
	result := FindFirstNonSolvable(parsedData.Field, int16(len(parsedData.CorruptionOrder)), width, height)
	sourcePoint := parsedData.CorruptionOrder[result]
	return Point{X: sourcePoint % width, Y: sourcePoint / width}
}

var Part2 func(in io.Reader) Point

var Part1 func(in io.Reader) int16

const fieldDim = 71

func init() {
	Part1 = func(in io.Reader) int16 {
		return Part1Base(in, 1024, fieldDim, fieldDim)
	}
	Part2 = func(in io.Reader) Point {
		return Part2Base(in, fieldDim, fieldDim)
	}
}
