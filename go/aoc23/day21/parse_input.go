package day21

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func parseInput(r io.Reader) (*parsedInput, error) {

	scanner := bufio.NewScanner(r)

	ret := parsedInput{
		data:   nil,
		Height: 0,
	}
	for scanner.Scan() {
		line := scanner.Bytes()
		if ret.data == nil {
			ret.data = make([]byte, 0, len(line)*len(line))
			ret.Width = len(line)
		}
		for idx, c := range line {
			switch c {
			case cellEmpty, cellStone:
				ret.data = append(ret.data, c)
			case cellStart:
				ret.data = append(ret.data, cellEmpty)
				ret.StartX = idx
				ret.StartY = ret.Height
			default:
				return nil, aoc_utils.NewUnexpectedInputError(c)
			}
		}
		ret.Height++
	}
	// sanity check, that all rows are filled
	if len(ret.data) != ret.Width*ret.Height {
		return nil, aoc_utils.NewParseError("malformed input: not all rows have the same length", nil)
	}
	return &ret, nil
}
