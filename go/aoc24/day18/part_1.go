package day18

import (
	"bufio"
	"io"
	"math"

	"github.com/donmahallem/aoc/aoc_utils"
)

const CELL_CORRUPTED int = -1

type Point = aoc_utils.Point[int16]

// As you always walk top-left to right-bottom primarly use those first
var DIRS_ALL [4]Point = [4]Point{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1}}

type Field = []int16

type ParseResult struct {
	Field           Field
	CorruptionOrder []int16
}

func ParseInput(in io.Reader, width, height int16) ParseResult {
	s := bufio.NewScanner(in)
	field := make(Field, width*height)
	order := make([]int16, 0, width*height)

	for pointIdx := int16(1); s.Scan(); pointIdx++ {
		line := s.Bytes()
		var currentX, currentY int16 = 0, 0
		target := &currentX
		for _, c := range line {
			if c == ',' {
				target = &currentY
			} else if c >= '0' && c <= '9' {
				*target = (*target * 10) + int16(c-'0')
			}
		}
		idx := int16(currentY*width + currentX)
		field[idx] = pointIdx
		order = append(order, idx)
	}
	return ParseResult{Field: field, CorruptionOrder: order}
}

type PathNode struct {
	X, Y, Steps int16
}

func FindShortestPath(field Field, stepsTaken, fieldWidth, fieldHeight int16) int16 {
	totalCells := fieldWidth * fieldHeight
	if totalCells <= 0 {
		return math.MaxInt16
	}

	startIdx := 0
	endIdx := totalCells - 1
	if (field[startIdx] > 0 && field[startIdx] <= stepsTaken) ||
		(field[endIdx] > 0 && field[endIdx] <= stepsTaken) {
		return math.MaxInt16
	}

	visited := make([]bool, totalCells)
	queue := make([]PathNode, 0, totalCells)
	queue = append(queue, PathNode{0, 0, 0})
	visited[startIdx] = true

	for head := 0; head < len(queue); head++ {
		currentNode := queue[head]

		for _, dir := range DIRS_ALL {
			nextX := currentNode.X + dir.X
			nextY := currentNode.Y + dir.Y
			if nextX < 0 || nextY < 0 || nextX >= fieldWidth || nextY >= fieldHeight {
				continue
			}

			idx := nextY*fieldWidth + nextX
			if visited[idx] {
				continue
			}

			cellValue := field[idx]
			if cellValue > 0 && cellValue <= stepsTaken {
				continue
			}

			nextLen := currentNode.Steps + 1
			if nextX == fieldWidth-1 && nextY == fieldHeight-1 {
				return nextLen
			}

			visited[idx] = true
			queue = append(queue, PathNode{X: nextX, Y: nextY, Steps: nextLen})
		}
	}
	return math.MaxInt16
}

func Part1Base(in io.Reader, maxSteps, width, height int16) int16 {
	parsedInput := ParseInput(in, width, height)
	return FindShortestPath(parsedInput.Field, maxSteps, width, height)
}
