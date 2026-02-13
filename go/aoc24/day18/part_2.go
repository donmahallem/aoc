package day18

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func IsPathAvailable(field Field, pointIdx, fieldWidth, fieldHeight int16) bool {

	totalCells := fieldWidth * fieldHeight
	startIdx := int16(0)
	endIdx := int16(totalCells - 1)
	if (field[startIdx] > 0 && field[startIdx] <= pointIdx) ||
		(field[endIdx] > 0 && field[endIdx] <= pointIdx) {
		return false
	}

	queue := make([]int16, 0, 128)
	queue = append(queue, startIdx)

	visited := make([]bool, totalCells)
	visited[0] = true

	for len(queue) > 0 {
		positionIdx := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		currentPositionX := positionIdx % fieldWidth
		currentPositionY := positionIdx / fieldWidth

		for _, checkDir := range DIRS_ALL {
			nextX := currentPositionX + checkDir.X
			nextY := currentPositionY + checkDir.Y
			if nextX < 0 || nextY < 0 || nextX >= fieldWidth || nextY >= fieldHeight {
				continue
			}

			nextPointIdx := nextY*fieldWidth + nextX
			nextIdxInt := int(nextPointIdx)
			if visited[nextIdxInt] {
				continue
			}

			cellValue := field[nextPointIdx]
			if cellValue > 0 && cellValue <= pointIdx {
				continue
			}

			if nextPointIdx == endIdx {
				return true
			}

			visited[nextIdxInt] = true
			queue = append(queue, nextPointIdx)
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

func Part2Base(in io.Reader, width, height int16) ([]int16, error) {
	parsedData, err := ParseInput(in, width, height)
	if err != nil {
		return []int16{}, err
	}
	if len(parsedData.CorruptionOrder) == 0 {
		return []int16{}, aoc_utils.NewParseError("Expected input to be atleast one", nil)
	}
	result := FindFirstNonSolvable(parsedData.Field, int16(len(parsedData.CorruptionOrder)), width, height)
	sourcePoint := parsedData.CorruptionOrder[result]
	return []int16{sourcePoint % width, sourcePoint / width}, nil
}

var Part2 func(in io.Reader) ([]int16, error)

var Part1 func(in io.Reader) (int16, error)

const fieldDim = 71

func init() {
	Part1 = func(in io.Reader) (int16, error) {
		return Part1Base(in, 1024, fieldDim, fieldDim)
	}
	Part2 = func(in io.Reader) ([]int16, error) {
		return Part2Base(in, fieldDim, fieldDim)
	}
}
