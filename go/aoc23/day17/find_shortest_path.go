package day17

func findShortestPath(f *field, minStraight, maxStraight uint8) uint32 {
	if f.Width == 0 || f.Height == 0 {
		return 0
	}
	if f.Width == 1 && f.Height == 1 {
		return 0
	}

	maxCost := ^uint32(0)

	width, height := f.Width, f.Height
	targetX, targetY := width-1, height-1

	directions := [4]struct {
		dx, dy int16
	}{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	stateIndex := func(x, y int16, dir, steps uint8) uint32 {
		cellIdx := uint32(y*width + x)
		return ((cellIdx*4 + uint32(dir)) * uint32(maxStraight)) + uint32(steps-1)
	}

	totalStates := uint32(width*height) * uint32(len(directions)) * uint32(maxStraight)
	best := make([]uint32, totalStates)
	for i := range best {
		best[i] = maxCost
	}

	cellCosts := f.Cells
	queue := make([]step, 0, len(cellCosts))

	// Loop over directions
	for dirIdx := uint8(0); dirIdx < uint8(len(directions)); dirIdx++ {
		dir := directions[dirIdx]
		nx, ny := dir.dx, dir.dy

		if nx < 0 || nx >= width || ny < 0 || ny >= height {
			continue
		}
		idxCell := int(ny)*int(width) + int(nx)
		if idxCell < 0 || idxCell >= len(cellCosts) {
			continue
		}
		cost := cellCosts[idxCell]
		idx := stateIndex(nx, ny, dirIdx, 1)
		if cost >= best[idx] {
			continue
		}
		best[idx] = cost
		queue = append(queue, step{
			x:     nx,
			y:     ny,
			dir:   dirIdx,
			steps: 1,
			cost:  cost,
		})
	}

	bestTarget := maxCost

	for head := 0; head < len(queue); head++ {
		current := &queue[head]

		if current.cost >= bestTarget {
			continue
		}

		stateIdx := stateIndex(current.x, current.y, current.dir, current.steps)
		if stored := best[stateIdx]; stored < current.cost {
			continue
		}

		if current.x == targetX && current.y == targetY {
			if current.steps >= minStraight {
				if current.cost < bestTarget {
					bestTarget = current.cost
				}
			}
			continue
		}

		if current.steps < minStraight {
			// Must continue in the same direction
			dir := directions[current.dir]
			nx := current.x + dir.dx
			ny := current.y + dir.dy

			if nx < 0 || nx >= width || ny < 0 || ny >= height {
				continue
			}

			newCost := current.cost + cellCosts[ny*width+nx]
			if newCost >= bestTarget {
				continue
			}

			newStateIdx := stateIndex(nx, ny, current.dir, current.steps+1)

			if best[newStateIdx] <= newCost {
				continue
			}
			best[newStateIdx] = newCost

			queue = append(queue, step{
				x:     nx,
				y:     ny,
				dir:   current.dir,
				steps: current.steps + 1,
				cost:  newCost,
			})
			continue
		}

		// Loop over directions
		for nextDirIdx := uint8(0); nextDirIdx < uint8(len(directions)); nextDirIdx++ {
			dir := directions[nextDirIdx]

			if nextDirIdx == (current.dir+2)&3 {
				continue
			}

			nx := current.x + dir.dx
			ny := current.y + dir.dy

			if nx < 0 || nx >= width || ny < 0 || ny >= height {
				continue
			}

			var steps uint8
			if nextDirIdx == current.dir {
				if current.steps >= maxStraight {
					continue
				}
				steps = current.steps + 1
			} else {
				steps = 1
			}

			newCost := current.cost + cellCosts[ny*width+nx]
			if newCost >= bestTarget {
				continue
			}

			newStateIdx := stateIndex(nx, ny, nextDirIdx, steps)

			if best[newStateIdx] <= newCost {
				continue
			}
			best[newStateIdx] = newCost

			queue = append(queue, step{
				x:     nx,
				y:     ny,
				dir:   nextDirIdx,
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
