package day18

import (
	"bufio"
	"container/heap"
	"io"
	"math"
	"sync"

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

type QueueItem struct {
	Idx      int16
	Priority int16
}

type PriorityQueue []*QueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*QueueItem))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func FindShortestPath(field Field, stepsTaken, fieldWidth, fieldHeight int16) int16 {
	var itemPool = sync.Pool{
		New: func() interface{} {
			return &QueueItem{}
		},
	}

	totalCells := fieldWidth * fieldHeight
	if totalCells <= 0 {
		return math.MaxInt16
	}

	startIdx := int16(0)
	endIdx := totalCells - 1
	if (field[startIdx] > 0 && field[startIdx] <= stepsTaken) ||
		(field[endIdx] > 0 && field[endIdx] <= stepsTaken) {
		return math.MaxInt16
	}

	targetX := fieldWidth - 1
	targetY := fieldHeight - 1

	// Store steps taken to reach each cell
	steps := make([]int16, totalCells)
	for i := range steps {
		steps[i] = math.MaxInt16
	}
	steps[startIdx] = 0

	pq := make(PriorityQueue, 0, totalCells)
	heap.Init(&pq)

	// Manhattan distance heuristic
	heuristic := func(x, y int16) int16 {
		return (targetX - x) + (targetY - y)
	}

	item := itemPool.Get().(*QueueItem)
	item.Idx = startIdx
	item.Priority = heuristic(0, 0)
	heap.Push(&pq, item)

	for pq.Len() > 0 {
		currentItem := heap.Pop(&pq).(*QueueItem)
		currentIdx := currentItem.Idx
		currentSteps := steps[currentIdx]

		// Return item to pool after use
		itemPool.Put(currentItem)

		currentX := currentIdx % fieldWidth
		currentY := currentIdx / fieldWidth

		if currentX == targetX && currentY == targetY {
			return currentSteps
		}

		for _, dir := range DIRS_ALL {
			nextX := currentX + dir.X
			nextY := currentY + dir.Y
			if nextX < 0 || nextY < 0 || nextX >= fieldWidth || nextY >= fieldHeight {
				continue
			}

			nextIdx := nextY*fieldWidth + nextX
			cellValue := field[nextIdx]
			if cellValue > 0 && cellValue <= stepsTaken {
				continue
			}

			nextSteps := currentSteps + 1
			if nextSteps < steps[nextIdx] {
				steps[nextIdx] = nextSteps
				priority := nextSteps + heuristic(nextX, nextY)

				nextItem := itemPool.Get().(*QueueItem)
				nextItem.Idx = nextIdx
				nextItem.Priority = priority
				heap.Push(&pq, nextItem)
			}
		}
	}
	return math.MaxInt16
}

func Part1Base(in io.Reader, maxSteps, width, height int16) int16 {
	parsedInput := ParseInput(in, width, height)
	return FindShortestPath(parsedInput.Field, maxSteps, width, height)
}
