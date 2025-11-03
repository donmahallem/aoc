package day15

import (
	"bufio"
	"fmt"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

const (
	CELL_WALL      uint8 = 1
	CELL_EMPTY     uint8 = 0
	CELL_PLAYER    uint8 = 2
	CELL_BOX       uint8 = 3
	CELL_BOX_LEFT  uint8 = 4
	CELL_BOX_RIGHT uint8 = 5
)

type Field = [][]uint8
type Move = aoc_utils.Point[int16]
type Player = aoc_utils.Point[int16]

func translateFieldRow(b *[]byte, doubleWide *bool) (*[]uint8, *int16) {
	size := len(*b)
	if *doubleWide {
		size *= 2
	}
	d := make([]uint8, size)
	player := int16(-1)
	var startX int
	for i := range len(*b) {
		startX = i
		if *doubleWide {
			startX *= 2
		}
		switch (*b)[i] {
		case '#':
			d[startX] = CELL_WALL
			if *doubleWide {
				d[startX+1] = CELL_WALL
			}
		case 'O':
			if *doubleWide {
				d[startX] = CELL_BOX_LEFT
				d[startX+1] = CELL_BOX_RIGHT
			} else {
				d[startX] = CELL_BOX
			}
		case '@':
			player = int16(startX)
		}
	}
	return &d, &player
}

func TranslateMovements(a *[]byte) *[]Move {
	movements := make([]Move, 0, len(*a))
	for i := range len(*a) {
		movements = append(movements, translateMovement((*a)[i]))
	}
	return &movements
}
func translateMovement(a byte) Move {
	if a == '<' {
		return Move{X: -1, Y: 0}
	} else if a == '>' {
		return Move{X: 1, Y: 0}
	} else if a == 'v' {
		return Move{X: 0, Y: 1}
	} else if a == '^' {
		return Move{X: 0, Y: -1}
	}
	panic("Illegal character provided")
}

func PrintField(field *Field, player *Player) {
	height, width := int16(len(*field)), int16(len((*field)[0]))
	for y := range height {
		for x := range width {
			if player.X == x && player.Y == y {
				fmt.Print("@")
				continue
			}
			switch (*field)[y][x] {
			case CELL_BOX:
				fmt.Print("O")
			case CELL_BOX_LEFT:
				fmt.Print("[")
			case CELL_BOX_RIGHT:
				fmt.Print("]")
			case CELL_WALL:
				fmt.Print("#")
			default:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
func ParseInput(in io.Reader, doubleWide bool) (Field, Player, []Move) {
	field := make(Field, 0)
	moves := make([]Move, 0)
	var player Player
	s := bufio.NewScanner(in)
	insideField := true
	for s.Scan() {
		line := s.Bytes()
		if len(line) == 0 {
			insideField = false
			continue
		} else if insideField {
			row, tmpPlayer := translateFieldRow(&line, &doubleWide)
			if *tmpPlayer >= 0 {
				player.Y = int16(len(field))
				player.X = *tmpPlayer
			}
			field = append(field, *row)
		} else {
			moves = append(moves, (*TranslateMovements(&line))...)
		}
	}
	return field, player, moves
}

func FindNextEmptyCellOffset(field *Field, player *Player, move *Move) (int16, bool) {
	curX, curY := (*player).X, (*player).Y
	for i := int16(0); ; i++ {
		curX += (*move).X
		curY += (*move).Y
		if (*field)[curY][curX] == CELL_EMPTY {
			return i, true
		} else if (*field)[curY][curX] == CELL_BOX {
			continue
		} else {
			return 0, false
		}
	}
}

func Walk(field *Field, player *Player, moves *[]Move, moveIdx uint) {
	if len(*moves) == int(moveIdx) {
		return
	}
	nextY := (*player).Y + (*moves)[moveIdx].Y
	nextX := (*player).X + (*moves)[moveIdx].X
	if (*field)[nextY][nextX] == CELL_EMPTY {
		(*player).Y = nextY
		(*player).X = nextX
	} else if (*field)[nextY][nextX] == CELL_BOX {
		offset, found := FindNextEmptyCellOffset(field, player, &(*moves)[moveIdx])
		if found {
			(*player).Y = nextY
			(*player).X = nextX
			(*field)[nextY][nextX] = CELL_EMPTY
			(*field)[nextY+(offset*(*moves)[moveIdx].Y)][nextX+(offset*(*moves)[moveIdx].X)] = CELL_BOX
		}
	}
	Walk(field, player, moves, moveIdx+1)
}

func CalculateGpsScore(field *Field) int {
	result := 0
	for y := range len(*field) {
		for x := range len((*field)[0]) {
			if (*field)[y][x] == CELL_BOX || (*field)[y][x] == CELL_BOX_LEFT {
				result += y*100 + x
			}
		}
	}
	return result
}

func Part1(in io.Reader) int {
	field, player, moves := ParseInput(in, false)
	Walk(&field, &player, &moves, 0)
	return CalculateGpsScore(&field)
}
