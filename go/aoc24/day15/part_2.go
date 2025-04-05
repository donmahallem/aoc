package day15

import (
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

type Candidate = aoc_utils.Point[int16]

func TryMoveHorizontal(field *Field, position *Player, move *Move) bool {
	var walk func(int16) bool
	walk = func(x int16) bool {
		nextX := x + move.X
		if (*field)[position.Y][nextX] == CELL_EMPTY {
			(*field)[position.Y][nextX] = (*field)[position.Y][x]
			return true
		} else if (*field)[position.Y][nextX] == CELL_WALL {
			return false
		} else if (*field)[position.Y][nextX] == CELL_BOX_LEFT || (*field)[position.Y][nextX] == CELL_BOX_RIGHT {
			if walk(nextX) {
				(*field)[position.Y][nextX] = (*field)[position.Y][x]
				return true
			}
		}
		return false
	}
	if walk(position.X) {
		(*field)[position.Y][position.X] = CELL_EMPTY
		position.X += move.X
		return true
	}
	return false
}

func IsMovableVertical(field *Field, position *Player, move *Move) bool {
	if (*field)[position.Y][position.X] == CELL_BOX_RIGHT {
		return IsMovableVertical(field, &Player{X: position.X - 1, Y: position.Y}, move)
	}
	status := true
	nextY := (*move).Y + (*position).Y
	nextValueLeft := (*field)[nextY][(*position).X]
	nextValueRight := (*field)[nextY][(*position).X+1]
	if nextValueLeft == CELL_WALL || nextValueRight == CELL_WALL {
		return false
	} else if nextValueLeft == CELL_EMPTY && nextValueRight == CELL_EMPTY {
		return true
	} else if nextValueLeft == CELL_BOX_LEFT {
		nextPos := Player{Y: nextY, X: (*position).X}
		status = status && IsMovableVertical(field, &nextPos, move)
		if status {
			return true
		} else {
			return false
		}
	}
	if nextValueLeft == CELL_BOX_RIGHT {
		nextPos := Player{Y: nextY, X: (*position).X - 1}
		status = status && IsMovableVertical(field, &nextPos, move)
		if !status {
			return false
		}
	}
	if nextValueRight == CELL_BOX_LEFT {
		nextPos := Player{Y: nextY, X: (*position).X + 1}
		status = status && IsMovableVertical(field, &nextPos, move)
		if !status {
			return false
		}
	}
	return status
}
func MoveBoxesVertically(field *Field, position *Player, move *Move) {
	if (*field)[position.Y][position.X] == CELL_BOX_RIGHT {
		MoveBoxesVertically(field, &Player{X: position.X - 1, Y: position.Y}, move)
		return
	}
	nextY := (*move).Y + (*position).Y
	nextValueLeft := (*field)[nextY][position.X]
	nextValueRight := (*field)[nextY][position.X+1]
	if nextValueLeft == CELL_BOX_LEFT {
		nextPos := Player{Y: nextY, X: position.X}
		MoveBoxesVertically(field, &nextPos, move)
	}
	if nextValueLeft == CELL_BOX_RIGHT {
		nextPos := Player{Y: nextY, X: position.X - 1}
		MoveBoxesVertically(field, &nextPos, move)
	}
	if nextValueRight == CELL_BOX_LEFT {
		nextPos := Player{Y: nextY, X: position.X + 1}
		MoveBoxesVertically(field, &nextPos, move)
	}
	(*field)[position.Y][position.X] = CELL_EMPTY
	(*field)[position.Y][position.X+1] = CELL_EMPTY
	(*field)[nextY][position.X] = CELL_BOX_LEFT
	(*field)[nextY][position.X+1] = CELL_BOX_RIGHT
}

func TryMoveVertical(field *Field, player *Player, move *Move) bool {
	nextPos := Player{X: player.X, Y: player.Y + move.Y}
	switch (*field)[nextPos.Y][nextPos.X] {
	case CELL_EMPTY:
		player.Y = nextPos.Y
		return true
	case CELL_WALL:
		return false
	}
	ok := IsMovableVertical(field, &nextPos, move)
	if ok {
		MoveBoxesVertically(field, &nextPos, move)
		player.Y += move.Y
	}
	return ok
}

func WalkWideBoxes(field *Field, player *Player, moves *[]Move) {
	for moveIdx := range len(*moves) {
		if (*moves)[moveIdx].X != 0 {
			TryMoveHorizontal(field, player, &(*moves)[moveIdx])
		} else if (*moves)[moveIdx].Y != 0 {
			TryMoveVertical(field, player, &(*moves)[moveIdx])
		}
	}
}

func Part2(in io.Reader) int {
	field, player, moves := ParseInput(in, true)
	WalkWideBoxes(&field, &player, &moves)
	return CalculateGpsScore(&field)
}
