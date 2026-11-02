package day02

import (
	"strconv"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type Dir int

const (
	Forward Dir = iota
	Down
	Up
)

func (d Dir) String() string {
	switch d {
	case Forward:
		return "forward"
	case Down:
		return "down"
	case Up:
		return "up"
	default:
		return "unknown"
	}
}

func parseBytesToInt(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, strconv.ErrSyntax
	}
	val := 0
	for _, ch := range b {
		if ch < '0' || ch > '9' {
			return 0, strconv.ErrSyntax
		}
		val = val*10 + int(ch-'0')
	}
	return val, nil
}
func parseLine(line []byte) (Dir, int, error) {
	n := len(line)
	if n < 3 {
		return 0, 0, aoc_utils.NewParseError("Short line", strconv.ErrSyntax)
	}

	switch line[0] {
	case 'f': // "forward"
		if n >= 8 && line[1] == 'o' && line[2] == 'r' && line[3] == 'w' && line[4] == 'a' && line[5] == 'r' && line[6] == 'd' {
			val, err := parseBytesToInt(line[8:])
			if err != nil {
				return 0, 0, err
			}
			return Forward, val, nil
		}
	case 'd': // "down"
		if n >= 5 && line[1] == 'o' && line[2] == 'w' && line[3] == 'n' {
			val, err := parseBytesToInt(line[5:])
			if err != nil {
				return 0, 0, err
			}
			return Down, val, nil
		}
	case 'u': // "up"
		if n >= 4 && line[1] == 'p' {
			val, err := parseBytesToInt(line[3:])
			if err != nil {
				return 0, 0, err
			}
			return Up, val, nil
		}
	}

	return 0, 0, aoc_utils.NewParseError("Invalid line", strconv.ErrSyntax)
}
