package day15

import (
	"io"
)

func tryMoveHorizontal(field *field, position *player, move *move) bool {
	height := len(*field)
	if int(position.Y) < 0 || int(position.Y) >= height {
		return false
	}
	width := len((*field)[0])

	// Find the end of the chain of boxes
	currX := position.X + move.X
	for {
		if int(currX) < 0 || int(currX) >= width {
			return false
		}
		cell := (*field)[position.Y][currX]
		if cell == CELL_WALL {
			return false
		}
		if cell == CELL_EMPTY {
			break
		}
		// It's a box part, continue
		currX += move.X
	}

	// Shift everything from the empty spot back to the player
	for x := currX; x != position.X; x -= move.X {
		if int(x-move.X) < 0 || int(x-move.X) >= width {
			return false
		}
		(*field)[position.Y][x] = (*field)[position.Y][x-move.X]
	}

	(*field)[position.Y][position.X] = CELL_EMPTY
	position.X += move.X
	return true
}

func tryMoveVertical(field *field, pl *player, move *move, boxesToMove *[]player, queue *[]player, visited []int, width int, gen int) bool {
	height := len(*field)
	// bounds checks
	nextY := int(pl.Y) + int(move.Y)
	if nextY < 0 || nextY >= height || int(pl.X) < 0 || int(pl.X) >= width {
		return false
	}
	nextCell := (*field)[nextY][pl.X]

	if nextCell == CELL_WALL {
		return false
	}
	if nextCell == CELL_EMPTY {
		pl.Y = int16(nextY)
		return true
	}

	// It's a box. Collect all affected boxes.
	// Clear buffers
	*boxesToMove = (*boxesToMove)[:0]
	*queue = (*queue)[:0]

	// Add the initial box(es)
	if nextCell == CELL_BOX_LEFT {
		p := player{X: pl.X, Y: int16(nextY)}
		*queue = append(*queue, p)
		visited[int(p.Y)*width+int(p.X)] = gen
	} else { // CELL_BOX_RIGHT
		// if the right half is first, ensure there is a left half to its left
		if int(pl.X)-1 < 0 {
			return false
		}
		p := player{X: pl.X - 1, Y: int16(nextY)}
		*queue = append(*queue, p)
		visited[int(p.Y)*width+int(p.X)] = gen
	}

	// BFS to find all connected boxes
	head := 0
	for head < len(*queue) {
		curr := (*queue)[head]
		head++
		*boxesToMove = append(*boxesToMove, curr)

		// Check what this box pushes
		nextBoxY := curr.Y + move.Y

		// Check left side
		leftVal := (*field)[nextBoxY][curr.X]
		if leftVal == CELL_WALL {
			return false
		}
		if leftVal == CELL_BOX_LEFT {
			p := player{X: curr.X, Y: nextBoxY}
			idx := int(p.Y)*width + int(p.X)
			if visited[idx] != gen {
				visited[idx] = gen
				*queue = append(*queue, p)
			}
		} else if leftVal == CELL_BOX_RIGHT {
			p := player{X: curr.X - 1, Y: nextBoxY}
			idx := int(p.Y)*width + int(p.X)
			if visited[idx] != gen {
				visited[idx] = gen
				*queue = append(*queue, p)
			}
		}

		// Check right side
		rightVal := (*field)[nextBoxY][curr.X+1]
		if rightVal == CELL_WALL {
			return false
		}
		if rightVal == CELL_BOX_LEFT {
			p := player{X: curr.X + 1, Y: nextBoxY}
			idx := int(p.Y)*width + int(p.X)
			if visited[idx] != gen {
				visited[idx] = gen
				*queue = append(*queue, p)
			}
		}
	}

	// If we reached here, all boxes can move.
	// Move them in reverse order (deepest first)
	for i := len(*boxesToMove) - 1; i >= 0; i-- {
		b := (*boxesToMove)[i]
		// Clear old position
		(*field)[b.Y][b.X] = CELL_EMPTY
		(*field)[b.Y][b.X+1] = CELL_EMPTY
		// Set new position
		(*field)[b.Y+move.Y][b.X] = CELL_BOX_LEFT
		(*field)[b.Y+move.Y][b.X+1] = CELL_BOX_RIGHT
	}

	pl.Y = int16(nextY)
	return true
}

func walkWideBoxes(field *field, pl *player, moves *[]move) {
	boxesToMove := make([]player, 0, 64)
	queue := make([]player, 0, 64)

	height := len(*field)
	width := len((*field)[0])
	visited := make([]int, height*width)
	gen := 0

	for moveIdx := range *moves {
		if (*moves)[moveIdx].X != 0 {
			tryMoveHorizontal(field, pl, &(*moves)[moveIdx])
		} else if (*moves)[moveIdx].Y != 0 {
			gen++
			tryMoveVertical(field, pl, &(*moves)[moveIdx], &boxesToMove, &queue, visited, width, gen)
		}
	}
}

func Part2(in io.Reader) (int, error) {
	data, err := parseInput(in, true)
	if err != nil {
		return 0, err
	}
	walkWideBoxes(&data.Field, &data.Player, &data.Movements)
	return calculateGpsScore(&data.Field), nil
}
