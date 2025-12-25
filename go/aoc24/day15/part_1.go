package day15

import (
	"io"
)

func findNextEmptyCellOffset(field *field, player *player, move *move) (int16, bool) {
	height := len(*field)
	if height == 0 {
		return 0, false
	}
	width := len((*field)[0])
	curX, curY := int((*player).X), int((*player).Y)
	for i := int16(0); ; i++ {
		curX += int((*move).X)
		curY += int((*move).Y)
		// bounds check
		if curY < 0 || curY >= height || curX < 0 || curX >= width {
			return 0, false
		}
		cell := (*field)[curY][curX]
		if cell == CELL_EMPTY {
			return i, true
		} else if cell == CELL_BOX {
			continue
		} else {
			return 0, false
		}
	}
}

func walk(field *field, player *player, moves *[]move) {
	height := len(*field)
	if height == 0 {
		return
	}
	width := len((*field)[0])
	for _, m := range *moves {
		nextY := int((*player).Y + m.Y)
		nextX := int((*player).X + m.X)
		// bounds check
		if nextY < 0 || nextY >= height || nextX < 0 || nextX >= width {
			// treat out-of-range as wall / ignore movement
			continue
		}
		cell := (*field)[nextY][nextX]

		if cell == CELL_EMPTY {
			(*player).Y = int16(nextY)
			(*player).X = int16(nextX)
		} else if cell == CELL_BOX {
			offset, found := findNextEmptyCellOffset(field, player, &m)
			if found {
				(*player).Y = int16(nextY)
				(*player).X = int16(nextX)
				(*field)[nextY][nextX] = CELL_EMPTY
				endY := nextY + int(offset*m.Y)
				endX := nextX + int(offset*m.X)
				if endY >= 0 && endY < height && endX >= 0 && endX < width {
					(*field)[endY][endX] = CELL_BOX
				}
			}
		}
	}
}

func calculateGpsScore(field *field) int {
	result := 0
	height := len(*field)
	if height == 0 {
		return 0
	}
	width := len((*field)[0])
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if (*field)[y][x] == CELL_BOX || (*field)[y][x] == CELL_BOX_LEFT {
				result += y*100 + x
			}
		}
	}
	return result
}

func Part1(in io.Reader) (int, error) {
	data, err := parseInput(in, false)
	if err != nil {
		return 0, err
	}
	walk(&data.Field, &data.Player, &data.Movements)
	return calculateGpsScore(&data.Field), nil
}
