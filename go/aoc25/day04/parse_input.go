package day04

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type field []byte
type inputData struct {
	field      field
	width, row int
}

func parseInput(in io.Reader) (*inputData, error) {
	s := bufio.NewScanner(in)
	var g field
	rows := 0
	width := 0

	for s.Scan() {
		line := s.Bytes()
		if width == 0 {
			width = len(line)
			g = make(field, 0, width*width)
		}
		for i := range len(line) {
			switch line[i] {
			case '@':
				g = append(g, 1)
			case '.':
				g = append(g, 0)
			default:
				return nil, aoc_utils.NewUnexpectedInputError(line[i])
			}
		}
		rows++
	}
	// sanity check
	if len(g) != width*rows {
		return nil, aoc_utils.NewParseError("Malformed input data", nil)
	}
	return &inputData{field: g, width: width, row: rows}, nil
}
