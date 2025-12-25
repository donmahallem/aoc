package day15

import (
	"bufio"
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

type field = [][]uint8
type move = aoc_utils.Point[int16]
type player = aoc_utils.Point[int16]

type inputData struct {
	Field     field
	Player    player
	Movements []move
}

func parseField(scanner *bufio.Scanner, doubleWide bool, field *field, pl *player) error {
	fieldWidth := -1
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 && len(*field) > 0 {
			// Field finished; ensure a player position was found while parsing the field
			if pl.X < 0 || pl.Y < 0 {
				return aoc_utils.NewParseError("no player parsed in field", nil)
			}
			return nil
		}
		expectedLength := len(line)
		if doubleWide {
			expectedLength *= 2
		}
		parsedLine := make([]uint8, 0, expectedLength)
		if len(line) > 0 && (line[0] != '#' || line[len(line)-1] != '#') {
			return aoc_utils.NewParseError("field line does not start and end with wall", nil)
		}
		for c := range line {
			switch line[c] {
			case '#':
				parsedLine = append(parsedLine, CELL_WALL)
				if doubleWide {
					parsedLine = append(parsedLine, CELL_WALL)
				}
			case 'O':
				if doubleWide {
					parsedLine = append(parsedLine, CELL_BOX_LEFT)
					parsedLine = append(parsedLine, CELL_BOX_RIGHT)
				} else {
					parsedLine = append(parsedLine, CELL_BOX)
				}
			case '@':
				// Only one player allowed
				if pl.X >= 0 || pl.Y >= 0 {
					return aoc_utils.NewParseError("multiple player positions", nil)
				}
				if doubleWide {
					*pl = player{X: int16(len(parsedLine)), Y: int16(len(*field))}
					parsedLine = append(parsedLine, CELL_EMPTY, CELL_EMPTY)
				} else {
					*pl = player{X: int16(len(parsedLine)), Y: int16(len(*field))}
					parsedLine = append(parsedLine, CELL_EMPTY)
				}
			case '.':
				if doubleWide {
					parsedLine = append(parsedLine, CELL_EMPTY, CELL_EMPTY)
				} else {
					parsedLine = append(parsedLine, CELL_EMPTY)
				}
			default:
				return aoc_utils.NewUnexpectedInputError(line[c])
			}
		}

		if fieldWidth == -1 {
			fieldWidth = len(parsedLine)
		} else if fieldWidth != len(parsedLine) {
			return aoc_utils.NewParseError("Inconsistent line widths", nil)
		}
		*field = append(*field, parsedLine)
	}
	// EOF reached: if we parsed a field, ensure it had a player
	if len(*field) > 0 {
		if pl.X < 0 || pl.Y < 0 {
			return aoc_utils.NewParseError("no player parsed in field", nil)
		}
	}
	if fieldWidth < 3 || len(*field) < 3 {
		return aoc_utils.NewParseError("field too small", nil)
	}
	// validate top and bottom walls
	for x := 0; x < fieldWidth; x++ {
		if (*field)[0][x] != CELL_WALL {
			return aoc_utils.NewParseError("top wall incomplete", nil)
		}
		if (*field)[len(*field)-1][x] != CELL_WALL {
			return aoc_utils.NewParseError("bottom wall incomplete", nil)
		}
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil

}

func parseMovements(scanner *bufio.Scanner, moves *[]move) error {
	for scanner.Scan() {
		line := scanner.Bytes()
		for c := range line {
			switch line[c] {
			case '<':
				*moves = append(*moves, move{X: -1, Y: 0})
			case '>':
				*moves = append(*moves, move{X: 1, Y: 0})
			case 'v':
				*moves = append(*moves, move{X: 0, Y: 1})
			case '^':
				*moves = append(*moves, move{X: 0, Y: -1})
			default:
				return aoc_utils.NewUnexpectedInputError(line[c])
			}
		}
	}
	return nil
}

func parseInput(in io.Reader, doubleWide bool) (*inputData, error) {
	field := make(field, 0)
	moves := make([]move, 0)
	var player player = player{X: -1, Y: -1}
	s := bufio.NewScanner(in)
	err := parseField(s, doubleWide, &field, &player)
	if err != nil {
		return nil, err
	}
	err = parseMovements(s, &moves)
	if err != nil {
		return nil, err
	}
	if len(moves) == 0 {
		return nil, aoc_utils.NewParseError("no movements parsed", nil)
	}
	if len(field) == 0 {
		return nil, aoc_utils.NewParseError("no field parsed", nil)
	}
	if len(field[0]) == 0 {
		return nil, aoc_utils.NewParseError("empty field parsed", nil)
	}
	if player.X < 0 || player.Y < 0 {
		return nil, aoc_utils.NewParseError("no player parsed", nil)
	}
	if int(player.X) < 0 || int(player.Y) < 0 || int(player.Y) >= len(field) || int(player.X) >= len(field[0]) {
		return nil, aoc_utils.NewParseError("player out of bounds", nil)
	}
	return &inputData{
		Field:     field,
		Player:    player,
		Movements: moves,
	}, nil
}
