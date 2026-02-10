package day17

import (
	"bufio"
	_ "embed"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

//go:embed sample.txt
var sample string

type cell = uint32

type step struct {
	x, y  int16
	dir   uint8
	steps uint8
	cost  uint32
}

type field struct {
	Cells         []cell
	Width, Height int16
}

func parseInput(r io.Reader) (*field, error) {
	scanner := bufio.NewScanner(r)

	f := field{
		Cells:  make([]cell, 0, 64),
		Width:  0,
		Height: 0,
	}
	for scanner.Scan() {
		line := scanner.Bytes()
		if f.Width == 0 {
			f.Width = int16(len(line))
		} else if f.Width != int16(len(line)) {
			return nil, aoc_utils.NewParseError("inconsistent line lengths in input", nil)
		}
		f.Height++

		for _, char := range line {
			if char < '0' || char > '9' {
				return nil, aoc_utils.NewUnexpectedInputError(char)
			}
			f.Cells = append(f.Cells, uint32(char-'0'))
		}
	}
	if f.Width == 0 || f.Height == 0 {
		return nil, aoc_utils.NewParseError("empty field", nil)
	}
	if len(f.Cells) != int(f.Width)*int(f.Height) {
		return nil, aoc_utils.NewParseError("inconsistent field size", nil)
	}
	return &f, nil
}
