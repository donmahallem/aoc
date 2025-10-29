package day21

import (
	"bufio"
	"io"
)

const (
	cellEmpty byte = '.'
	cellStone byte = '#'
	cellStart byte = 'S'
)

type parsedInput struct {
	data   []byte
	Width  int
	Height int
	StartX int
	StartY int
}

func ParseInput(r io.Reader) parsedInput {

	scanner := bufio.NewScanner(r)

	ret := parsedInput{
		data:   nil,
		Height: 0,
	}
	for scanner.Scan() {
		line := scanner.Bytes()
		if ret.data == nil {
			ret.data = make([]byte, 0, len(line)*len(line))
			ret.Width = len(line)
		}
		for idx, c := range line {
			switch c {
			case cellEmpty, cellStone:
				ret.data = append(ret.data, c)
			case cellStart:
				ret.data = append(ret.data, cellEmpty)
				ret.StartX = idx
				ret.StartY = ret.Height
			}
		}
		ret.Height++
	}
	return ret
}

var directions = [][2]int{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

type queueItem struct {
	x     int
	y     int
	depth int
}

func CountVisited(inp *parsedInput, maxDepth int) int {

	queue := make([]queueItem, 0)
	visited := make(map[int]byte)
	startIdx := inp.StartY*inp.Width + inp.StartX
	queue = append(queue, queueItem{
		x:     inp.StartX,
		y:     inp.StartY,
		depth: 0,
	})
	visited[startIdx] = 2
	for len(queue) > 0 {
		currentQueueItem := queue[0]
		queue = queue[1:]
		x := currentQueueItem.x
		y := currentQueueItem.y
		if currentQueueItem.depth >= maxDepth {
			continue
		}
		for _, dir := range directions {
			newX := x + dir[0]
			newY := y + dir[1]
			if newX < 0 || newX >= inp.Width || newY < 0 || newY >= inp.Height {
				continue
			}
			newIdx := newY*inp.Width + newX
			if inp.data[newIdx] == cellStone {
				continue
			}
			newDepth := currentQueueItem.depth + 1
			cellValue := byte(1)
			if newDepth%2 == 0 {
				cellValue = 2
			}
			if visited[newIdx]&cellValue != 0 {
				continue
			}
			visited[newIdx] |= cellValue
			if newDepth < maxDepth {
				queue = append(queue, queueItem{
					x:     newX,
					y:     newY,
					depth: newDepth,
				})
			}
		}
	}
	targetParity := byte(1)
	if maxDepth%2 == 0 {
		targetParity = 2
	}
	counter := 0
	for _, v := range visited {
		if v&targetParity != 0 {
			counter++
		}
	}
	return counter
}
func Part1(in io.Reader) int {
	inp := ParseInput(in)
	return CountVisited(&inp, 64)
}
