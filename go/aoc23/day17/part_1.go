package day17

import (
	"bufio"
	"io"
)

type Cell = uint

type Field struct {
	Cells         []Cell
	Width, Height int16
}

func ParseInputPart1(r io.Reader) Field {
	scanner := bufio.NewScanner(r)

	field := Field{
		Cells:  make([]Cell, 0, 64),
		Width:  0,
		Height: 0,
	}
	for scanner.Scan() {
		line := scanner.Bytes()
		if field.Width == 0 {
			field.Width = int16(len(line))
		}
		field.Height++
		for idx := range line {
			field.Cells = append(field.Cells, uint(line[idx]-'0'))
		}
	}
	return field
}

func findShortestPath(field Field) uint {
	if field.Width == 0 || field.Height == 0 {
		return 0
	}
	if field.Width == 1 && field.Height == 1 {
		return 0
	}

	const maxStraight = uint8(3)
	maxCost := ^uint(0)

	width, height := field.Width, field.Height
	widthInt, heightInt := int(width), int(height)
	targetX, targetY := width-1, height-1

	directions := [4]struct {
		dx, dy int16
	}{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	stateIndex := func(x, y int16, dir, steps uint8) int16 {
		cellIdx := y*width + x
		return ((cellIdx*4 + int16(dir)) * int16(maxStraight)) + int16(steps-1)
	}

	totalStates := widthInt * heightInt * len(directions) * int(maxStraight)
	best := make([]uint, totalStates)
	for i := range best {
		best[i] = maxCost
	}

	cellCosts := field.Cells
	queue := make([]step, 0, len(cellCosts))

	for dirIdx, dir := range directions {
		nx, ny := dir.dx, dir.dy
		if nx < 0 || nx >= width || ny < 0 || ny >= height {
			continue
		}
		cost := cellCosts[ny*width+nx]
		idx := stateIndex(nx, ny, uint8(dirIdx), 1)
		if cost >= best[idx] {
			continue
		}
		best[idx] = cost
		queue = append(queue, step{
			x:     nx,
			y:     ny,
			dir:   uint8(dirIdx),
			steps: 1,
			cost:  cost,
		})
	}

	bestTarget := maxCost

	for head := 0; head < len(queue); head++ {
		current := queue[head]

		if current.cost >= bestTarget {
			continue
		}

		stateIdx := stateIndex(current.x, current.y, current.dir, current.steps)
		if stored := best[stateIdx]; stored < current.cost {
			continue
		}

		if current.x == targetX && current.y == targetY {
			if current.cost < bestTarget {
				bestTarget = current.cost
			}
			continue
		}

		for nextDirIdx, dir := range directions {
			if nextDirIdx == (int(current.dir)+2)&3 {
				continue
			}
			nx := current.x + dir.dx
			ny := current.y + dir.dy
			if nx < 0 || nx >= width || ny < 0 || ny >= height {
				continue
			}

			var steps uint8
			if nextDirIdx == int(current.dir) {
				if current.steps >= maxStraight {
					continue
				}
				steps = current.steps + 1
			} else {
				steps = 1
			}

			newCost := current.cost + uint(cellCosts[ny*width+nx])
			if newCost >= bestTarget {
				continue
			}

			newStateIdx := stateIndex(nx, ny, uint8(nextDirIdx), steps)
			if best[newStateIdx] <= newCost {
				continue
			}
			best[newStateIdx] = newCost

			queue = append(queue, step{
				x:     int16(nx),
				y:     int16(ny),
				dir:   uint8(nextDirIdx),
				steps: steps,
				cost:  newCost,
			})
		}
	}

	if bestTarget == maxCost {
		return 0
	}
	return bestTarget
}

func Part1(in io.Reader) uint {
	start := ParseInputPart1(in)
	return findShortestPath(start)
}
