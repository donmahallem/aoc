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
	// sets the first line as expected width for the field. -1 indicates it hasn't been set yet.
	expectedLength := -1
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 && len(*field) > 0 {
			return validateField(*field, *pl)
		}
		if expectedLength == -1 {
			expectedLength = len(line)
			if doubleWide {
				expectedLength *= 2
			}
		}
		currentLineLength := len(line)
		if doubleWide {
			currentLineLength *= 2
		}
		if currentLineLength != expectedLength {
			return aoc_utils.NewParseError("inconsistent line length", nil)
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

		*field = append(*field, parsedLine)
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return validateField(*field, *pl)
}

// validateField ensures the bounds checks in the solvers can rely on a walled field
func validateField(field field, pl player) error {
	if len(field) < 3 || len(field[0]) < 3 {
		return aoc_utils.NewParseError("field too small", nil)
	}
	if pl.X < 0 || pl.Y < 0 {
		return aoc_utils.NewParseError("no player parsed in field", nil)
	}
	for x := range field[0] {
		if field[0][x] != CELL_WALL {
			return aoc_utils.NewParseError("top wall incomplete", nil)
		}
		if field[len(field)-1][x] != CELL_WALL {
			return aoc_utils.NewParseError("bottom wall incomplete", nil)
		}
	}
	return nil
}

func parseMovements(scanner *bufio.Scanner, moves *[]move) error {
	buf := *moves
	for scanner.Scan() {
		line := scanner.Bytes()
		for c := range line {
			switch line[c] {
			case '<':
				buf = append(buf, move{X: -1, Y: 0})
			case '>':
				buf = append(buf, move{X: 1, Y: 0})
			case 'v':
				buf = append(buf, move{X: 0, Y: 1})
			case '^':
				buf = append(buf, move{X: 0, Y: -1})
			default:
				return aoc_utils.NewUnexpectedInputError(line[c])
			}
		}
	}
	*moves = buf
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
