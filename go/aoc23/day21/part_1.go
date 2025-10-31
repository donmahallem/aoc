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

func CountVisited(inp *parsedInput, maxDepth int, finite bool) int {
	queue := []queueItem{{x: inp.StartX, y: inp.StartY, depth: 0}}
	visited := make(map[[2]int]byte, len(inp.data))
	startKey := [2]int{inp.StartX, inp.StartY}
	visited[startKey] = 2

	width := inp.Width
	height := inp.Height

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.depth >= maxDepth {
			continue
		}

		for _, dir := range directions {
			newX := current.x + dir[0]
			newY := current.y + dir[1]

			var transposedX, transposedY int
			if finite {
				// part 1 finite grid
				if newX < 0 || newX >= width || newY < 0 || newY >= height {
					continue
				}
				transposedX, transposedY = newX, newY
			} else {
				// part 2 infinite grid
				// wraps coordinates to grid
				transposedX = newX % width
				if transposedX < 0 {
					transposedX += width
				}
				transposedY = newY % height
				if transposedY < 0 {
					transposedY += height
				}
			}

			if inp.data[transposedY*width+transposedX] == cellStone {
				continue
			}

			nextDepth := current.depth + 1
			visitByte := byte(1)
			if nextDepth%2 == 0 {
				visitByte = 2
			}

			key := [2]int{newX, newY}
			if visited[key]&visitByte != 0 {
				continue
			}
			visited[key] |= visitByte

			if nextDepth < maxDepth {
				queue = append(queue, queueItem{x: newX, y: newY, depth: nextDepth})
			}
		}
	}

	targetParity := byte(1)
	if maxDepth%2 == 0 {
		targetParity = 2
	}

	count := 0
	for _, val := range visited {
		if val&targetParity != 0 {
			count++
		}
	}
	return count
}

func CountVisitedInfinite(inp *parsedInput, maxDepth int) int {
	return CountVisited(inp, maxDepth, false)
}

func Part1(in io.Reader) int {
	inp := ParseInput(in)
	return CountVisited(&inp, 64, true)
}
