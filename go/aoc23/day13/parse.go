package day13

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type block struct {
	Rows []int
	Cols []int
}

func parseInput(r io.Reader) ([]block, error) {
	scanner := bufio.NewScanner(r)

	blocks := make([]block, 0, 4)
	b := block{
		Rows: make([]int, 0, 16),
		Cols: nil,
	}

	currentHeight := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			currentHeight = 0
			if len(b.Rows) > 0 {
				blocks = append(blocks, b)
			}
			b = block{
				Rows: make([]int, 0, 16),
				Cols: nil,
			}
			continue
		}
		if b.Cols == nil {
			b.Cols = make([]int, len(line))
		}
		if len(line) != len(b.Cols) {
			return nil, aoc_utils.NewParseError("inconsistent line widths in block", nil)
		}
		currentLine := 0
		for idx, c := range line {
			switch c {
			case '#':
				currentLine |= 1 << idx
				b.Cols[idx] |= 1 << currentHeight
			case '.':
				// do nothing
			default:
				return nil, aoc_utils.NewUnexpectedInputError(c)
			}
		}
		b.Rows = append(b.Rows, currentLine)
		currentHeight++
	}
	if currentHeight > 0 {
		blocks = append(blocks, b)
	}
	return blocks, nil
}
